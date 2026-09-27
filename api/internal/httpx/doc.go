// Package httpx holds transport plumbing shared by every handler: the router,
// middleware, request decoding, and response/error writing.
//
// EXERCISE (M2)
// -------------
// router.go — Go 1.22+ ServeMux understands method+pattern routes, so you need
// no third-party router:
//
//	mux.HandleFunc("POST /users", h.Register)
//	mux.HandleFunc("GET /articles/{slug}", h.GetArticle)   // r.PathValue("slug")
//
// middleware.go — write these as func(http.Handler) http.Handler and compose
// them:
//   - RequestID: generate one, put it in the context and in a response header.
//   - Logger: method, path, status, duration, request id. You will need a
//     ResponseWriter wrapper to capture the status — write it and understand why.
//   - Recover: turn a panic into a 500 plus a logged stack, not a dead process.
//   - CORS: the Vite dev server proxies /api, so you can ignore this locally —
//     but deployed frontends hit you cross-origin. See
//     .specs/realworld/docs/.../backend/cors.md.
//   - Authenticate (required) and MaybeAuthenticate (optional). The spec has
//     endpoints of both kinds: GET /profiles/:username changes its `following`
//     field depending on whether a token was supplied.
//
// respond.go — one place that writes JSON, so Content-Type is always
// "application/json; charset=utf-8" as the spec demands, and one place that maps
// domain errors to status codes:
//
//	domain.ErrNotFound      -> 404
//	domain.ErrForbidden     -> 403   (authenticated, but not allowed)
//	missing/invalid token   -> 401
//	domain.ValidationError  -> 422 {"errors":{"field":["message"]}}
//
// Getting 401 vs 403 vs 404 right is not pedantry: it is a whole Hurl file
// (.specs/api/hurl/errors_authorization.hurl).
package httpx
