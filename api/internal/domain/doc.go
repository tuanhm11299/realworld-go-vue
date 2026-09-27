// Package domain holds the Conduit entities and the rules that govern them,
// expressed in plain Go with no knowledge of HTTP or SQL.
//
// EXERCISE (M1+)
// --------------
// Define the types the spec talks about — User, Profile, Article, Comment, Tag —
// and the operations that are genuinely about the domain rather than about
// transport or storage. For example:
//
//   - Article.Slug generation and uniqueness (the spec only requires a unique
//     string; see .specs/realworld/docs/.../backend/endpoints.md#update-article).
//   - Whether a given user may edit or delete a given article.
//   - Sentinel errors the layers above translate into status codes:
//     ErrNotFound, ErrForbidden, ErrConflict, and a ValidationError carrying
//     field -> messages so handlers can render the spec's 422 shape:
//     {"errors":{"body":["can't be empty"]}}
//
// The discipline to hold: this package imports neither net/http nor pgx. If you
// feel the urge to import one, the code you are writing belongs in
// internal/handler or internal/store instead. That constraint is the whole
// point — it is what makes the domain testable without a server or a database.
package domain
