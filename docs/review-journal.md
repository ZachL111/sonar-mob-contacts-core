# Review Journal

This journal records the domain cases that matter before widening the public API.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 127, lane `watch`
- `stress`: `sync drift`, score 100, lane `hold`
- `edge`: `local state`, score 201, lane `ship`
- `recovery`: `conflict cost`, score 192, lane `ship`
- `stale`: `form pressure`, score 134, lane `watch`

## Note

A future change should add new cases before it changes the scoring rule.
