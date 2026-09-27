// Package config turns the process environment into a validated, typed Config.
//
// EXERCISE (M0)
// -------------
// cmd/api/main.go currently reads os.Getenv inline. That is fine for two
// variables and awful for twenty. Build this package instead:
//
//   - Config struct: DatabaseURL, Port, JWTSecret, JWTTTL, Env (dev/prod),
//     LogLevel, CORSOrigins.
//   - Load() (Config, error) that reads the environment, applies defaults, and
//     returns *all* validation failures at once (a config error should tell you
//     everything that is wrong, not just the first thing).
//   - Refuse to start in prod with the development JWT secret.
//
// Then delete the env() helper from main.go and wire Load() in.
//
// Things worth thinking about while you do it:
//   - Why return a value rather than populate a package-level global?
//   - Where should defaults live: here, or in the Makefile/.env.example?
//   - How do you test this without mutating the real environment? (Hint: pass
//     in a func(string) string, or use t.Setenv.)
package config
