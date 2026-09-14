-- +goose Up
-- +goose StatementBegin

-- EXERCISE (M1): design the Conduit schema here.
--
-- This migration is intentionally empty. Applying it proves the goose pipeline
-- works end to end (`make migrate-up`, `make migrate-status`, `make migrate-down`)
-- without handing you the schema — designing it is the first real piece of work.
--
-- The spec needs to support, at minimum:
--
--   users          unique email, unique username, password hash, nullable bio,
--                  nullable image, timestamps
--   follows        who follows whom; a user must not follow themselves twice
--   articles       unique slug, title, description, body, author, timestamps
--   tags           the tag vocabulary
--   article_tags   many-to-many between articles and tags
--   favorites      which user favourited which article
--   comments       body, author, article, timestamps
--
-- Decisions to make deliberately, and to record in docs/decisions/ as ADRs:
--
--   * Keys: bigserial or uuid? What does each cost you at the API boundary?
--   * Emails: citext, or a lower(email) unique index? Case sensitivity here is
--     a real bug source.
--   * Deletes: ON DELETE CASCADE, or explicit deletes in a transaction?
--   * Timestamps: timestamptz, always. (Why not `timestamp`? Find out before
--     you type it.)
--   * Which columns need indexes for the feed and the filtered article list?
--     Write them when M5 shows you the query plan, not before — but predict
--     them now and check yourself later.
--
-- Prefer several small migrations over one enormous one: you will want to
-- practise rolling individual changes back.

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Every Up needs a Down that actually reverses it. If you cannot write the
-- Down, you have learned something important about the Up.

-- +goose StatementEnd
