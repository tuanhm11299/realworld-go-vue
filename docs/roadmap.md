# Beyond CRUD

Conduit is a CRUD app. Building one well is table stakes — it is not seniority.
What follows is the part that goes further: ten stages, each built **in this
codebase**, on a system you already understand, against a conformance suite that
already passes so you can tell the moment you break something.

That last point is the whole reason to do this here rather than in ten tutorial
projects. Real engineering is mostly *changing* a working system without
breaking it. You cannot practise that on a greenfield repo.

## How to use this

- **One stage at a time, roughly one per week.** They are ordered so that each
  gives you tools the next one needs — you cannot tune what you cannot measure,
  and you should not re-architect what you have no tests for.
- **Every stage ends with an artifact.** A flame graph, a dashboard, a benchmark
  diff, an ADR, a postmortem. "I read about it" is not a finish line. The
  *Prove it* line for each stage is the deliverable.
- **Run `make verify` before and after.** If the suite went from green to red,
  that is the lesson, not an interruption.
- **Write an ADR** (`docs/decisions/`) for anything you would have to defend in
  a review.

---

## Stage A · Testing depth

*Prerequisite for everything else. Do this one first.*

**Why:** you currently have someone else's end-to-end tests and almost none of
your own. E2E suites are slow, coarse, and tell you *that* something broke, not
where. You need a test pyramid before you start changing things underneath it.

**Build:**
- Table-driven unit tests for the parts with real logic: slug generation, the
  partial-update decoder, JWT parsing, the error-to-status mapping.
- Handler tests with `httptest` and a faked store (sqlc's `emit_interface` is
  already on for exactly this).
- Integration tests against a **real Postgres** via
  [testcontainers-go](https://golang.testcontainers.org/) — an in-memory fake
  would not catch the constraint violations you care about.
- `go test -fuzz` on the slug generator and the JWT parser.
- Frontend: Vitest component tests with [MSW](https://mswjs.io/) stubbing the API.
- Mutation testing with [gremlins](https://github.com/go-gremlins/gremlins) to
  find tests that assert nothing.

**Prove it:** delete a random line of business logic. A *unit* test must fail,
and fast. If only the e2e suite notices, your pyramid is upside down.

---

## Stage B · Observability

**Why:** in production you cannot attach a debugger. Everything you will ever
know about a live system arrives as logs, metrics and traces — and you have to
have put them there in advance.

**Build:**
- Structured `slog` everywhere, with a request id threaded through context and
  into every log line. Never log tokens, passwords or full request bodies.
- OpenTelemetry traces spanning Vue → Go → Postgres, so one page load is one
  waterfall. `otelpgx` instruments the driver.
- RED metrics (Rate, Errors, Duration) per endpoint, exported to Prometheus.
- A `docker compose --profile observability up` bringing Grafana, Tempo and Loki,
  with a dashboard committed to the repo.
- One SLO — say, "99% of feed requests under 300ms over 30 days" — and the error
  budget that follows from it.

**Prove it:** add an artificial 400ms sleep somewhere in the feed path, then find
it using *only* a trace waterfall. If you cannot, your spans are too coarse.

---

## Stage C · Performance

**Why:** performance work done by intuition is folklore. Done with a profiler it
is engineering, and the gap between the two is enormous.

**Build:**
- `net/http/pprof` behind a flag. Capture CPU, heap, block and mutex profiles
  under load, and read them as flame graphs.
- Load tests with [k6](https://k6.io) — a realistic mix, not one endpoint hammered.
- `EXPLAIN (ANALYZE, BUFFERS)` on the feed and the filtered article list. Fix the
  worst plan. Measure again.
- Tune the pgx pool: what happens at `max_conns=2`? At 200? Why is bigger not
  better?
- `benchstat` in CI comparing against the previous commit, to catch regressions.
- Frontend: Lighthouse CI, a bundle-size budget, route-level code splitting
  (`() => import(...)` in the router), and virtualised lists for long feeds.

**Prove it:** a before/after p99 number **with the profile that explained it**.
A speedup you cannot explain is luck, and luck does not generalise.

---

## Stage D · Concurrency and background work

**Why:** this is what Go is *for*, and Conduit as written barely uses it. It is
also where the genuinely hard bugs live — the ones that only appear under load
and never reproduce locally.

**Build:**
- `errgroup` to parallelise the independent queries behind an article page, with
  proper context cancellation. Measure whether it actually helped.
- A job queue on Postgres using `SELECT ... FOR UPDATE SKIP LOCKED` — no new
  infrastructure, and you will understand every line. Give it retries with
  backoff, a dead-letter table, and idempotency keys.
- Use it for something real: email on new follower, or a daily digest.
- The **outbox pattern**: write the job in the same transaction as the domain
  change, so you can never send a notification for a rollback.
- Extend the graceful shutdown you got in M0 to drain in-flight jobs.

**Prove it:** `kill -9` the worker mid-job. Nothing is lost, nothing is
double-processed. Then run the whole suite with `-race` and keep it clean.

---

## Stage E · Caching and data

**Why:** caching is easy to add and hard to invalidate, and most people learn it
by shipping a bug. Better to learn it where the bug costs nothing.

**Build:**
- Redis in front of `GET /api/tags` and the global feed. Measure first — you
  should be able to say what you saved.
- Then the hard half: invalidation. What happens when an article is deleted?
  Time-based, write-through, or explicit purge? Pick and defend.
- HTTP caching: `ETag` / `If-None-Match` on articles, so repeat loads are 304s.
- Replace offset pagination with cursor pagination on the feed, and explain why
  `OFFSET 10000` is slow.
- Full-text article search with a `tsvector` column and a GIN index.

**Prove it:** show the cache hit ratio and the 304s in your Stage B dashboard.
Then write the test that catches your stale-cache bug.

---

## Stage F · Real-time

**Why:** request/response is not the only shape a system can have, and the
problems of long-lived connections — backpressure, reconnection, fan-out — do not
exist in HTTP handlers.

**Build:**
- SSE (simpler) or WebSockets (bidirectional) for live comments on an article and
  notifications on a new follower.
- A Vue composable owning the connection: reconnect with backoff, resume from the
  last seen id, clean up on unmount.
- Fan-out across processes. A single server can hold connections in a map; two
  cannot. Postgres `LISTEN/NOTIFY` or Redis pub/sub bridges them.
- Backpressure: what happens when a client reads slower than you write?

**Prove it:** two browsers, one comment, no refresh. Then kill the server and
watch the client recover without losing messages.

---

## Stage G · Architecture

**Why:** you now have a codebase with real constraints and real tests, which is
the only honest place to evaluate an architectural idea. Patterns learned on
toy examples are cargo cult.

**Build:**
- Refactor to ports and adapters: the domain defines interfaces, `store` and
  `handler` implement them, and the domain imports neither `net/http` nor `pgx`.
  Judge it afterwards — is it better, or just more files?
- Separate domain errors from transport errors properly, if you have not.
- **Contract-first with OpenAPI.** Write `api/openapi.yml`, generate Go server
  interfaces with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
  and TypeScript types with
  [openapi-typescript](https://github.com/openapi-ts/openapi-typescript). One
  document, two languages, drift impossible. This is the single highest-value
  change on this list for day-to-day work.
- API versioning: add `/api/v2` with one breaking change, and support both.
- A CQRS read model for the feed — a denormalised table kept current by the
  outbox from Stage D. Now you have eventual consistency. How does the UI cope?
- Extract one service (notifications) over NATS or Kafka, then **write down what
  it cost you**: a network call that used to be a function call, a deploy
  ordering problem, a distributed trace to debug one request. Deciding *not* to
  do this in the end is a perfectly good outcome — as long as you can say why.

**Prove it:** an ADR per decision, including the ones you reverted. The reverted
ones are the valuable ones.

---

## Stage H · Security

**Why:** you have already met XSS in M9. That is one item on a long list, and
the rest do not come with a test suite that finds them for you.

**Build:**
- An OWASP Top 10 pass over your own code, written up finding by finding.
- Refresh tokens with rotation and reuse detection; move the access token to an
  httpOnly cookie and implement CSRF protection. Then write up the trade-off
  against the localStorage approach the spec chose — you will have built both.
- Rate limiting on auth endpoints, and a lockout policy that cannot be used to
  lock *other people* out.
- Tune argon2id properly: measure, target ~100ms, document the parameters.
- A Content Security Policy strict enough to have caught the M9 payloads.
- `govulncheck` and `npm audit` in CI; Dependabot on.
- Secrets: what happens today if `JWT_SECRET` leaks? Can you rotate it without
  logging everyone out? (Key ids in the JWT header.)
- An authorization model — roles, or a policy engine — now that ownership checks
  are scattered across handlers.

**Prove it:** for each finding, write the attack, then the fix, then the
regression test. In that order.

---

## Stage I · Delivery and operations

**Why:** software that is not deployed is a hobby. And the parts that are hard —
migrations against a live database, rollback, restoring from backup — are hard
precisely because nobody practises them.

**Build:**
- Multi-stage distroless Docker images for both halves. Get the Go image under
  20MB and understand every layer.
- **Expand/contract migrations**: rename a column without downtime, in three
  deploys. This is the single most useful operational skill on this list.
- GitHub Actions CD to a real host — Fly.io, Hetzner, a small AWS box — with
  Terraform for the infrastructure.
- Blue/green or canary deploys, with a rollback you have actually executed.
- Feature flags, so deploy and release stop being the same event.
- Automated backups **and a restore drill**.
- Inject a failure on purpose — kill the database mid-request, fill the disk,
  add 500ms of network latency — and see what your Stage B dashboards say.

**Prove it:** restore the database from a backup into a fresh environment, and
time it. Then do the column rename with the app serving traffic throughout.

---

## Stage J · Craft (continuous, not a phase)

Running underneath all of the above:

- **ADRs** for every decision you would have to defend. Include the rejected
  options; the reasoning is the artifact, not the conclusion.
- **Commits that explain why**, not what. The diff already says what.
- **Review your own PRs** before asking anyone else to. You will find things.
- **Read the source** when an API surprises you. `net/http`, `pgx` and Vue's
  reactivity core are all readable, and reading them is how "I use this" becomes
  "I understand this".
- **Write a postmortem** for the failure you injected in Stage I: timeline,
  impact, root cause, what would have caught it earlier. Blameless, specific.
- **Contribute upstream once.** A doc fix to `gothinkster/realworld`, a bug
  report with a reproduction, an issue on a library you used. Working in public
  is a skill with its own learning curve.
- **Teach one of these stages** to someone else. It is the fastest way to find
  the parts you only think you understand.

---

## A note on pace

Ten stages at one a week is a season's work, and doing three of them properly
beats skimming all ten. The order matters more than the speed: tests before
changes, measurement before optimisation, understanding before abstraction.

The goal was never Conduit. Conduit is just a system real enough to learn on.
