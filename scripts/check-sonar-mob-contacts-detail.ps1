$ErrorActionPreference = "Stop"
$root = Resolve-Path (Join-Path $PSScriptRoot "..")
Set-Location $root

$repo = "sonar-mob-contacts-core"
$language = "Go"
$category = "Mobile workflows"

foreach ($path in @("metadata/domain-review.json", "config/review-profile.json", "fixtures/domain_review.csv", "docs/field-notes.md", "docs/review-journal.md", "examples/sonar-mob-contacts-walkthrough.md")) {
  if (-not (Test-Path $path)) { throw "$repo missing $path" }
}

$metadata = Get-Content -Raw metadata/domain-review.json | ConvertFrom-Json
if ($metadata.repo -ne $repo) { throw "$repo metadata repo mismatch" }
if ($metadata.language -ne $language) { throw "$repo language mismatch" }
if ($metadata.category -ne $category) { throw "$repo category mismatch" }

$rows = Import-Csv fixtures/domain_review.csv
if ($rows.Count -ne 5) { throw "$repo expected five domain review rows" }
foreach ($row in $rows) {
  $score = ([int]$row.signal * 2) + [int]$row.slack + [int]$row.confidence - ([int]$row.drag * 3)
  if ($score -ne [int]$row.score) { throw "$repo score mismatch for $($row.name)" }
  if ($score -ge 140) { $lane = "ship" } elseif ($score -ge 105) { $lane = "watch" } else { $lane = "hold" }
  if ($lane -ne $row.lane) { throw "$repo lane mismatch for $($row.name)" }
}
$ship = @($rows | Where-Object { $_.lane -eq "ship" }).Count
$hold = @($rows | Where-Object { $_.lane -eq "hold" }).Count
if ($ship -ne 2) { throw "$repo ship lane count changed" }
if ($hold -ne 1) { throw "$repo hold lane count changed" }

$profile = Get-Content -Raw config/review-profile.json | ConvertFrom-Json
if ($profile.repo -ne $repo) { throw "$repo profile repo mismatch" }
if ($profile.case_count -ne $rows.Count) { throw "$repo profile case count mismatch" }
if ($profile.highest_case.name -ne "edge") { throw "$repo profile highest case changed" }
if ($profile.lowest_case.name -ne "stress") { throw "$repo profile lowest case changed" }

Write-Host "$repo detail check passed"
