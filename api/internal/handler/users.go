package handler

import "net/http"

// Register handles POST /api/users.
//
// No auth. Required fields: email, username, password (nested under "user").
// Returns a User: {"user":{email,token,username,bio,image}} — bio and image are
// null until set. Duplicate email or username, or a missing field, is a 422
// {"errors":{...}}.
//
// Tests: .specs/api/hurl/auth.hurl, .specs/api/hurl/errors_auth.hurl
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// Login handles POST /api/users/login.
//
// No auth. Required: email, password. Returns a User with a fresh token.
// Wrong password and unknown email must be indistinguishable to the caller —
// both 422 — or you have leaked which emails are registered.
//
// Tests: .specs/api/hurl/auth.hurl, .specs/api/hurl/errors_auth.hurl
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// CurrentUser handles GET /api/user.
//
// Auth required. Returns the authenticated User, including a token (the spec
// expects one on every User payload — decide whether you re-issue or echo).
//
// Tests: .specs/api/hurl/auth.hurl
func (h *Handler) CurrentUser(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// UpdateUser handles PUT /api/user.
//
// Auth required. Accepted fields: email, username, password, image, bio — all
// optional, and this is where partial updates get subtle. The upstream suite
// distinguishes three cases per field:
//
//	field absent       -> leave unchanged
//	field: ""          -> normalise an empty bio/image to null
//	field: null        -> accept for nullable fields
//
// encoding/json cannot tell "absent" from "zero value" with a plain string
// field. Use *string (or json.RawMessage, or a custom Optional[T]) and know why
// you chose it. This is the single most instructive endpoint in the spec.
//
// Tests: .specs/api/hurl/auth.hurl (cases 06-09), errors_auth.hurl
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) { todo(w, r) }
