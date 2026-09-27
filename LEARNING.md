# The syllabus

This repository is a course with an automatic grader. The grader is the upstream
RealWorld conformance suite; the course is ten milestones that take you from an
empty schema to a complete, hardened application.

The rule that makes it work: **do not read a reference implementation.** Dozens
exist. Reading one converts an exercise into a transcription task, and you will
feel like you understood something you did not. Read the spec, read the failing
test, write the code.

---

## How a milestone works

1. **Pick the next milestone** below and open `docs/milestones/Mx.md`.
2. **Run its tests and watch them fail.** Read the failure before writing code —
   a Hurl failure tells you the request, the expected response and what it got.
   ```bash
   .specs/api/run-api-tests-hurl.sh .specs/api/hurl/auth.hurl   # one backend file
   npx playwright test .specs/e2e/auth.spec.ts                  # one frontend file
   ```
3. **Write the smallest thing that makes the next assertion pass.** Not the
   whole milestone. The next assertion.
4. **Make it green, then make it good.** Working first, then naming, structure,
   error handling, tests of your own.
5. **Answer the milestone's questions in writing.** In a commit message, a note,
   an ADR under `docs/decisions/` — somewhere. This is the step people skip, and
   it is the step that turns "I made the test pass" into knowledge you keep.
6. **Commit per milestone, at least.** Your git history should read like a story.

If you are stuck for more than about 30 minutes, you are missing a concept
rather than a line of code. Go read about the concept.

---

## Milestones

| # | Milestone | Green when | Guide |
|---|---|---|---|
| **M0** | Walking skeleton | `/healthz`, `/readyz`, `make verify` runs all red | [M0](docs/milestones/M0.md) |
| **M1** | Schema and data layer | `make migrate-up && make sqlc` generate compiling Go | [M1](docs/milestones/M1.md) |
| **M2** | Authentication | `auth.hurl`, `errors_auth.hurl` | [M2](docs/milestones/M2.md) |
| **M3** | Profiles and following | `profiles.hurl`, `errors_profiles.hurl` | [M3](docs/milestones/M3.md) |
| **M4** | Articles and tags | `articles.hurl`, `errors_articles.hurl`, `tags.hurl` | [M4](docs/milestones/M4.md) |
| **M5** | Favorites, feed, pagination | `favorites.hurl`, `feed.hurl`, `pagination.hurl` | [M5](docs/milestones/M5.md) |
| **M6** | Comments and authorization | `comments.hurl`, `errors_comments.hurl`, `errors_authorization.hurl` | [M6](docs/milestones/M6.md) |
| **M7** | Vue shell | `health.spec.ts`, `auth.spec.ts`, `navigation.spec.ts` | [M7](docs/milestones/M7.md) |
| **M8** | Feature parity | `articles`, `comments`, `social`, `settings`, `url-navigation`, `null-fields` | [M8](docs/milestones/M8.md) |
| **M9** | Hardening | `error-handling`, `xss-security`, `user-fetch-errors`; CI conformance required | [M9](docs/milestones/M9.md) |

M0–M6 is Go. M7–M9 is Vue.

### Two ways through

**Backend first (M0 → M9 in order).** One language at a time, and the frontend
is easy once the API is real. Best if Go is the less familiar half.

**Full-stack loop (M0, M1, M2, M7, then M3–M6, M8, M9).** Do authentication end
to end — Postgres, Go, JWT, Pinia, a login form that actually logs in — before
going wide. Slower to a green backend, but you feel the whole system early and
you hit the interesting questions (where does the token live? who validates
what?) on day three instead of week three.

Pick deliberately. Do not drift between them.

---

## Working habits worth building now

- **Read the error before you change anything.** Guess-and-check is the slowest
  possible way to learn, and it is seductive because it sometimes works.
- **One concern per commit**, with a message that says *why*.
- **Write the test that would have caught the bug** — every time, not sometimes.
- **Keep `make lint` clean.** A linter you ignore is a linter you should delete.
- **Record decisions as ADRs** in `docs/decisions/`. Six weeks from now you will
  not remember why you chose `bigserial` over `uuid`, and you will re-litigate it.
- **Finish what you start.** A half-migrated codebase is worse than either end
  state, and getting comfortable with that is a career-limiting habit.

---

## When Conduit works

Conduit is a CRUD app. Building one well is table stakes; it is not seniority.
**[docs/roadmap.md](docs/roadmap.md)** is the part that goes further — ten stages
covering testing depth, observability, performance, concurrency, caching,
real-time, architecture, security, and delivery. Each one is built *in this
codebase*, on a system you understand, against tests that already pass so you
can tell when you break something.

That is the real curriculum. Conduit is how you earn a system worth doing it to.
