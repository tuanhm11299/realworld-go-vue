package handler

import "net/http"

// GetProfile handles GET /api/profiles/{username}.
//
// Auth OPTIONAL — and that is the lesson. Returns
// {"profile":{username,bio,image,following}} where `following` is false for an
// anonymous caller and reflects the real relationship for an authenticated one.
// Unknown username is 404.
//
// Tests: .specs/api/hurl/profiles.hurl, errors_profiles.hurl
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// FollowUser handles POST /api/profiles/{username}/follow.
//
// Auth required, no body. Returns the Profile with following=true. Following
// someone you already follow must succeed rather than error — make the write
// idempotent (ON CONFLICT DO NOTHING) instead of checking first and racing.
//
// Tests: .specs/api/hurl/profiles.hurl
func (h *Handler) FollowUser(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// UnfollowUser handles DELETE /api/profiles/{username}/follow.
//
// Auth required. Returns the Profile with following=false. Also idempotent.
//
// Tests: .specs/api/hurl/profiles.hurl
func (h *Handler) UnfollowUser(w http.ResponseWriter, r *http.Request) { todo(w, r) }
