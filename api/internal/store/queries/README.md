# sqlc queries

One `.sql` file per aggregate: `users.sql`, `profiles.sql`, `articles.sql`,
`comments.sql`, `tags.sql`, `favorites.sql`.

Each query carries an annotation telling sqlc what to generate:

```sql
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListArticlesByTag :many
SELECT * FROM articles a
JOIN article_tags at ON at.article_id = a.id
JOIN tags t ON t.id = at.tag_id
WHERE t.name = $1
ORDER BY a.created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteArticle :exec
DELETE FROM articles WHERE slug = $1;
```

`:one`, `:many`, `:exec`, `:execrows`, `:copyfrom` — pick deliberately.

Run `make sqlc` after every change. Until you have written both a schema
(M1) and at least one query, `sqlc generate` will report that it found
nothing — that is expected, not a broken setup.
