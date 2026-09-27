# 1. Record architecture decisions

- **Status:** Accepted
- **Date:** 2026-09-14

## Context

This project exists to learn, and most of what is worth learning is in the
*reasoning*, not the code. Six weeks after choosing `bigserial` over `uuid` you
will not remember why, and you will re-litigate it — or worse, silently reverse
it in one corner of the codebase.

A decision record is a few paragraphs written while the trade-off is still fresh.
It is also the closest thing to a design review when you are working alone.

## Decision

Record significant decisions as numbered Markdown files in this directory, using
the template below. "Significant" means: something you would have to defend in a
code review, or something a future reader would reasonably ask "why?" about.

Do not record trivia. Do not write one per commit.

## Consequences

- Reviewing your own decisions becomes possible.
- Reversing one is cheap and honest: add a new ADR that supersedes the old, and
  mark the old one `Superseded by ADR-00xx`. Never edit history to look right.
- Writing the "Alternatives considered" section will occasionally change your
  mind before you write any code. That is the point.

---

## Template

```markdown
# N. Short title in the imperative

- **Status:** Proposed | Accepted | Superseded by ADR-00xx
- **Date:** YYYY-MM-DD

## Context
What forced a decision? What constraints were in play?

## Decision
What did you choose? State it plainly.

## Alternatives considered
What else was on the table, and what specifically made you reject it?

## Consequences
What gets easier? What gets harder? What have you now committed to?
```

## Decisions this project will ask you to make

Milestones M1–M9 each end with questions worth an ADR. The ones most likely to
matter later:

- Primary key type (M1)
- Email case-sensitivity strategy (M1)
- Cascade deletes vs explicit transactional deletes (M1)
- Password hashing algorithm and parameters (M2)
- How "field absent" vs "field null" is modelled in Go (M2)
- `favoritesCount` computed vs denormalised (M5)
- Where frontend feed state lives: URL, component, or store (M8)
- Token storage: localStorage vs httpOnly cookie (M7, revisited in Stage H)
