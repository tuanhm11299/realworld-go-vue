package handler

import "net/http"

// ListArticles handles GET /api/articles.
//
// Auth optional. Most recent first. Query params: tag, author, favorited,
// limit (default 20), offset (default 0) — they combine.
//
// Note: since 2024-08-16 list responses OMIT the article body for performance.
// Returns {"articles":[...],"articlesCount":N} where articlesCount is the total
// matching the filter, NOT the length of this page.
//
// The naive implementation runs one query for the articles and then, per
// article, one for tags, one for the author, one for the favourite count. That
// is the N+1 problem. Get it correct first, then measure it with EXPLAIN
// ANALYZE and fix it — that sequence is the whole point of M5.
//
// Tests: .specs/api/hurl/articles.hurl, pagination.hurl, errors_articles.hurl
func (h *Handler) ListArticles(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// FeedArticles handles GET /api/articles/feed.
//
// Auth REQUIRED. Articles by authors the caller follows, most recent first,
// with limit/offset. Watch your route registration order and patterns: "feed"
// must not be swallowed by the {slug} route.
//
// Tests: .specs/api/hurl/feed.hurl
func (h *Handler) FeedArticles(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// GetArticle handles GET /api/articles/{slug}.
//
// Auth optional. Returns a single Article — this one DOES include the body, and
// `favorited` depends on the caller.
//
// Tests: .specs/api/hurl/articles.hurl, errors_articles.hurl
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// CreateArticle handles POST /api/articles.
//
// Auth required. Required: title, description, body. Optional: tagList.
// Creating the article and its tag links is one unit of work — do it in a
// transaction, and make sure two articles with the same title still get
// distinct slugs.
//
// Tests: .specs/api/hurl/articles.hurl, errors_articles.hurl
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// UpdateArticle handles PUT /api/articles/{slug}.
//
// Auth required, author only (someone else's article is 403, not 404). All
// fields optional; changing the title changes the slug.
//
// The upstream suite pins tagList semantics precisely:
//
//	absent  -> tags preserved            (case 12)
//	[]      -> all tags removed          (cases 13-14)
//	null    -> rejected                  (case 15)
//
// Same *string / RawMessage problem as UpdateUser. Solve it once, reuse it.
//
// Tests: .specs/api/hurl/articles.hurl (cases 10-15), errors_authorization.hurl
func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// DeleteArticle handles DELETE /api/articles/{slug}.
//
// Auth required, author only. Decide what happens to comments, tag links and
// favourites: ON DELETE CASCADE in the schema, or explicit deletes in a
// transaction? Both are defensible — know why you picked yours.
//
// Tests: .specs/api/hurl/articles.hurl, errors_authorization.hurl
func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// FavoriteArticle handles POST /api/articles/{slug}/favorite.
//
// Auth required, no body. Returns the Article with favorited=true and an
// updated favoritesCount. Idempotent.
//
// Tests: .specs/api/hurl/favorites.hurl
func (h *Handler) FavoriteArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }

// UnfavoriteArticle handles DELETE /api/articles/{slug}/favorite.
//
// Auth required. Returns the Article with favorited=false. Idempotent.
//
// Tests: .specs/api/hurl/favorites.hurl
func (h *Handler) UnfavoriteArticle(w http.ResponseWriter, r *http.Request) { todo(w, r) }
