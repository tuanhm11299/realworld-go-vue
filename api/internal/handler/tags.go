package handler

import "net/http"

// ListTags handles GET /api/tags.
//
// No auth. Returns {"tags":["reactjs","angularjs",...]}.
//
// The simplest endpoint in the spec, and the best one to come back to in Stage E
// of docs/roadmap.md: it is read-mostly, identical for every caller, and on a
// real corpus it is a sequential scan over every tag link. Perfect cache bait —
// but only once you can *measure* that it is slow.
//
// Tests: .specs/api/hurl/tags.hurl
func (h *Handler) ListTags(w http.ResponseWriter, r *http.Request) { todo(w, r) }
