#!/usr/bin/env bash
# Fetch the upstream RealWorld conformance suites into .specs/ (gitignored).
#
# We deliberately do NOT vendor these into git: they are thousands of lines of
# upstream MIT-licensed code that would drown your own diffs, and they get
# updated. Syncing keeps attribution clean and the contract current.
#
#   .specs/realworld/  full shallow clone of gothinkster/realworld
#   .specs/e2e         -> realworld/specs/e2e   (Playwright suite + SELECTORS.md)
#   .specs/api         -> realworld/specs/api   (Hurl suite + runner)
#   web/public/styles.css  <- assets/theme/styles.css (Conduit Minimal CSS)
set -euo pipefail

REPO_URL="https://github.com/gothinkster/realworld"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SPECS="$ROOT/.specs"
CLONE="$SPECS/realworld"

mkdir -p "$SPECS"

if [ -d "$CLONE/.git" ]; then
  echo "==> Updating $CLONE"
  git -C "$CLONE" fetch --depth 1 origin HEAD
  git -C "$CLONE" reset --hard FETCH_HEAD
else
  echo "==> Cloning $REPO_URL"
  rm -rf "$CLONE"
  git clone --depth 1 "$REPO_URL" "$CLONE"
fi

ln -sfn "$CLONE/specs/e2e" "$SPECS/e2e"
ln -sfn "$CLONE/specs/api" "$SPECS/api"
chmod +x "$CLONE"/specs/api/*.sh "$CLONE"/specs/api/hurl/*.sh 2>/dev/null || true

# The frontend spec requires the shared Conduit stylesheet.
cp "$CLONE/assets/theme/styles.css" "$ROOT/web/public/styles.css"
cp "$CLONE/assets/media/default-avatar.svg" "$ROOT/web/public/default-avatar.svg" 2>/dev/null || true

cat <<SUMMARY

==> Synced.
    .specs/api/hurl/     $(ls "$CLONE/specs/api/hurl"/*.hurl | wc -l | tr -d ' ') Hurl files  (backend conformance)
    .specs/e2e/          $(ls "$CLONE/specs/e2e"/*.spec.ts | wc -l | tr -d ' ') Playwright specs (frontend conformance)
    .specs/e2e/SELECTORS.md  the DOM contract your Vue components must honour
    web/public/styles.css    Conduit Minimal CSS

SUMMARY
