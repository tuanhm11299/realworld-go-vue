package handler

import "net/http"

// AddComment handles POST /api/articles/{slug}/comments.
//
// Auth required. Required: body. Returns
// {"comment":{id,createdAt,updatedAt,body,author}} where author is a Profile.
//
// Tests: .specs/api/hurl/comments.hurl, errors_comments.hurl
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// ListComments handles GET /api/articles/{slug}/comments.
//
// Auth optional — each author's `following` flag reflects the caller.
// Returns {"comments":[...]}.
//
// Tests: .specs/api/hurl/comments.hurl
func (h *Handler) ListComments(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// DeleteComment handles DELETE /api/articles/{slug}/comments/{id}.
//
// Auth required, comment author only. Consider the awkward cases the suite
// checks: a comment id that does not exist, an id that exists but belongs to a
// different article, and a valid comment owned by someone else. They are 404,
// 404 and 403 respectively — and telling them apart is the exercise.
//
// Tests: .specs/api/hurl/comments.hurl, errors_comments.hurl,
// errors_authorization.hurl
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) { todo(w, r) }
