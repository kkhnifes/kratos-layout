# Refactor mise task scripts

Scope: the three Go programs under `.mise/tasks/`. `mise.toml` inline tasks
are explicitly out of scope and must not change.

## Goals

Applied to every changed file, in priority order:

1. **Efficient & stable** — no needless work, robust to re-execution, errors
   surface clearly.
2. **Reexecutable, override-or-abort** — interactive setup tasks show the
   current value and let the operator keep it (empty input = keep current) or
   override it. Aborting via empty input is the default, not a dead end.
3. **Modular & reusable** — shared logic lives once; no duplicated helpers
   across scripts.
4. **Self-readable code** — no comments unless genuinely necessary; names
   carry the meaning.

## Current state

- `.mise/tasks/setup/appname.go` — hard-aborts if `APP_NAME` already set; no
  override path.
- `.mise/tasks/setup/gopkg.go` — duplicates `miseGet`/`miseSet`/`readLine`
  from appname.go; re-runs unconditionally (rewrites module path every time).
- `.mise/tasks/build/docker.go` — standalone; no shared code; thin but lacks
  preflight checks (`docker buildx` availability, empty version guard).

## Design

### Shared package: `.mise/internal/mise/`

A small internal package the task scripts import via the repo module path
(they run as separate `main` packages but share the enclosing go.mod).

Exports:

- `Get(key string) (string, error)` — wraps `mise config get`.
- `Set(key, value string) error` — wraps `mise config set <key>=<value>`.
- `Prompt(label, current string) (value string, changed bool, err error)`
  - Prints `label [current: <current>]: ` (omits the bracketed part when
    `current == ""`).
  - Reads one line from stdin.
  - Trimmed empty input → returns `current, false, nil` (keep / abort).
  - Trimmed non-empty input → returns the new value, true, nil.

No other helpers here; anything used by only one script stays local to that
script.

### appname.go

- Read current `APP_NAME` via `mise.Get`.
- `mise.Prompt("Enter app name", current)`.
- Unchanged → log "kept <current>", exit 0.
- Changed → `mise.Set`, log "set to <new>", exit 0.

No hard abort when already set; the prompt *is* the override gate.

### gopkg.go

- Read current module from `go.mod` (keep local `readGoModule`).
- Read current `PKG_PATH` and `APP_NAME` and `GIT_PROVIDER` via `mise.Get`.
- `mise.Prompt("New path", pkgPath)`.
- Unchanged → log "kept <current>", exit 0.
- Changed → `mise.Set` PKG_PATH; compute new module
  `<provider>/<newPath>/<app>`; `replaceOnce` in go.mod; `replaceInTree`;
  `go mod tidy`. Keep existing local helpers (`readGoModule`, `replaceOnce`,
  `replaceInTree`).

### docker.go

Standalone (no shared pkg needed). Tighten:

- Preflight: `exec.LookPath("docker")`; if missing → clear fatal.
- Guard empty `git describe` / `vcsRef` (treat as error, don't build a
  `:tag`).
- Keep existing `gitOut` helper; keep `os.LookupEnv` checks; keep the
  `buildx` invocation shape.

No override prompt on docker (out of scope per decision).

## Execution plan

Two independent agent+reviewer loops, run in parallel:

- **Agent A — setup**: create `.mise/internal/mise/` package; rewrite
  `appname.go` and `gopkg.go` to use it and gain override-or-abort.
- **Agent B — build**: tighten `docker.go` with preflight + guards.

Each agent's loop:
1. Implement.
2. Run `go build ./...` from repo root; fix until clean.
3. Dispatch a bullish reviewer subagent that critiques the diff against the
   four goals above; loop fixes until the reviewer approves with no blocking
   issues.

Final review by the main session (me) over the combined diff.

## Out of scope

- `mise.toml` changes.
- Override/confirm behavior on non-interactive tasks (generate, lint, fmt,
  security, docker push).
- New third-party dependencies.