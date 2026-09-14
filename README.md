# realworld-go-vue

A **learning scaffold** for building [RealWorld (Conduit)](https://github.com/gothinkster/realworld)
with a Go backend and a Vue 3 frontend.

This repository is deliberately incomplete. The tooling, the database, the test
harness and the CI pipeline are done; **the application is not**. What you get
instead of a finished app is a precise, executable list of everything that is
missing, taken from the upstream RealWorld conformance suites:

- **139 Playwright tests** across 12 files — the frontend contract.
- **13 Hurl files** covering every endpoint plus five files of error and
  authorization edge cases — the backend contract.

Run `make verify`, watch it all go red, and turn it green one milestone at a
time. Start with **[LEARNING.md](LEARNING.md)**.

There is no need to submit anything to the RealWorld platform — both suites run
locally and in CI.

---

## Stack

| Layer | Choice | Why |
|---|---|---|
| HTTP | Go stdlib `net/http` (1.22+ method+pattern routing) | No framework between you and the fundamentals |
| Database | PostgreSQL 16 + `pgx/v5` | The real thing, no abstraction tax |
| Queries | [sqlc](https://sqlc.dev) | You write SQL; it generates type-safe Go |
| Migrations | [goose](https://github.com/pressly/goose) | Plain SQL, up and down |
| Frontend | Vue 3 `<script setup>` + TypeScript + Pinia + Vue Router | The mainstream Vue stack |
| Build | Vite | Fast, and Vitest comes with it |
| Conformance | Hurl (API) + Playwright (e2e) | Upstream's own suites — an objective finish line |

## Prerequisites

| Tool | Version | Notes |
|---|---|---|
| Go | 1.24+ | |
| Node | 22+ | |
| Docker | any recent | For Postgres; a local Postgres works too |
| [Hurl](https://hurl.dev/docs/installation.html) | **8.0+** | Earlier versions cannot parse the suite's `isList` predicate |

`make tools` installs the Go CLIs (goose, sqlc, golangci-lint, air). Hurl you
install yourself — it is a Rust binary, not a Go one.

## Quickstart

```bash
git clone https://github.com/tuanhm11299/realworld-go-vue
cd realworld-go-vue

make setup          # .env, Go CLIs, upstream specs, dependencies
make up             # Postgres on :5432, Adminer on :8081
make migrate-up     # applies the (empty) initial migration
make dev            # API on :8080, web on :5173
```

Then, in another terminal:

```bash
make verify         # both conformance suites — this is your to-do list
```

Expect **0% passing**. That is the correct starting state.

Sanity checks that *should* pass on day one:

```bash
curl localhost:8080/healthz    # {"status":"ok"}
curl localhost:8080/readyz     # {"status":"ready"} — proves the DB is wired
open http://localhost:5173     # every route renders what it must become
```

Run `make` with no arguments to see every target.

## Layout

```
api/                     Go backend
  cmd/api/main.go        the walking skeleton — the only file with real logic
  migrations/            goose SQL migrations          (M1)
  internal/config/       env -> typed Config           (M0)
  internal/httpx/        router, middleware, errors    (M2)
  internal/auth/         password hashing, JWT         (M2)
  internal/domain/       entities and rules            (M1+)
  internal/store/        sqlc-generated queries        (M1)
  internal/handler/      one handler per endpoint      (M2-M6)

web/                     Vue 3 SPA
  src/router/            the 8 spec routes             (M7)
  src/views/             one stub per route            (M7-M9)
  src/stores/auth.ts     Pinia auth store              (M7)
  src/api/client.ts      fetch wrapper                 (M7)
  src/types/api.ts       API payload types             (M7)

docs/milestones/         the syllabus, M0 through M9
docs/roadmap.md          where to go once Conduit works
docs/decisions/          ADRs — record what you chose and why

.specs/                  upstream conformance suites (gitignored, `make specs-sync`)
playwright.config.ts     runs the upstream e2e suite against this stack
```

Every stub carries a doc comment with its contract, the spec section it comes
from, and the test file that proves it. You should rarely need to go hunting.

## Continuous integration

Three jobs (`.github/workflows/ci.yml`):

- **Go** — build, vet, gofmt, golangci-lint, `go test -race`
- **Web** — lint, typecheck, Vitest, build
- **Conformance** — spins up Postgres, runs both upstream suites

The conformance job is `continue-on-error` on purpose: it reports how many
upstream tests pass rather than failing the build, so the number going up is
your progress bar. Make it required when you finish M9.

## Credits

The specification, the Hurl collection, the Playwright suite and `styles.css`
come from [gothinkster/realworld](https://github.com/gothinkster/realworld)
(MIT). They are fetched into `.specs/` rather than vendored, so this repository
contains only your own work.
