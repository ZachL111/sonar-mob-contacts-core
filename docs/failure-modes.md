# Failure Modes

For `sonar-mob-contacts-core`, I would look first for these mistakes:

- `form pressure` cases moving lanes without a matching threshold change.
- `local state` scoring higher after drag increases.
- Duplicate fixture ids hiding a stale golden row.
- README examples drifting away from the verifier.

The local checks are intentionally strict about these cases.
