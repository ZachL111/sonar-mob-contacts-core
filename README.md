# sonar-mob-contacts-core

`sonar-mob-contacts-core` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for contacts workflows, centered on incremental indexing, append-only fixtures, and checkpoint recovery checks.

## Why I Keep It Small

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how form pressure and local state should influence a review result.

## Sonar Mob Contacts Core Review Notes

`edge` and `stress` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Included Behavior

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/sonar-mob-contacts-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `local state` and `sync drift`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Internal Model

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Try It Locally

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Validation

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Scope

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
