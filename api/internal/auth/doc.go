// Package auth handles password hashing and JWT issuing/verification.
//
// EXERCISE (M2)
// -------------
// password.go:
//   - Hash(plain string) (string, error) and Verify(hash, plain string) error.
//   - Use argon2id (golang.org/x/crypto/argon2) or bcrypt. If argon2id, store
//     the parameters in the encoded hash so you can raise them later without
//     invalidating existing passwords.
//   - Verify must take the same time for a wrong password as for an unknown
//     user, or you have built a username oracle.
//
// jwt.go:
//   - Issue(userID, ttl) (string, error) and Parse(token) (userID, error).
//   - The spec's header is non-standard: `Authorization: Token jwt.token.here`
//     — note "Token", not "Bearer". .specs/api/hurl/errors_auth.hurl checks this.
//   - Reject tokens with alg=none and with an unexpected signing method. This
//     is the classic JWT vulnerability; write the test that proves you are safe.
//
// Questions to answer before moving on:
//   - What exactly is in your token, and what happens when that data changes?
//   - How would you revoke one today? (You can't — which is why Stage H of
//     docs/roadmap.md adds refresh tokens.)
package auth
