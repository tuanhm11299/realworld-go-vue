/**
 * TypeScript types for the Conduit API payloads.
 *
 * EXERCISE (M1/M7)
 * ----------------
 * Write these by reading the spec, not by guessing from your Go structs —
 * transcribing the contract by hand is how you notice its sharp edges:
 *
 *   User     { email, token, username, bio: string | null, image: string | null }
 *   Profile  { username, bio: string | null, image: string | null, following: boolean }
 *   Article  { slug, title, description, body, tagList: string[], createdAt,
 *              updatedAt, favorited, favoritesCount, author: Profile }
 *   Comment  { id: number, createdAt, updatedAt, body, author: Profile }
 *
 * Two traps the upstream tests care about:
 *   1. Every payload is WRAPPED: {"user": {...}}, {"articles": [...], "articlesCount": n}.
 *      Model the envelope, do not silently unwrap in three different places.
 *   2. List endpoints omit `body` from articles (since 2024-08-16). So the thing
 *      in a feed is NOT the same type as the thing on an article page. Encoding
 *      that in the type system — ArticlePreview vs Article — stops a whole class
 *      of undefined-at-runtime bugs.
 *
 * Stage G of docs/roadmap.md replaces this file with types generated from a
 * shared OpenAPI document, so the Go server and the Vue client can never drift.
 * Writing them by hand first is what makes that pay-off obvious.
 */

export {}
