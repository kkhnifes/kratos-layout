# Biz

Domain logic layer. Owns the domain object (DO), usecase, and repo interface.
Pure Go — no proto tags, no storage tags, no storage clients.

## Contents per resource (see `todo.go` for the pattern)

- **DO** (`type <Resource> struct`): pure domain model, no proto, no storage tags.
- **Repo interface** (`type <Resource>Repo interface`): the inversion seam
  `data` implements; declared here, never imported here.
- **Usecase** (`type <Resource>Usecase`): domain logic built on the repo
  interface. Validation and rules live here.
- **Typed errors** (`Err<Resource><Cause>`): built with
  `errors.NotFound` / `errors.BadRequest` plus the reason enum from
  `api/<domain>/<version>/error_reason.proto`.
- **`ListOption` helpers** — `ListFilter`, `ListOrderBy`, `ListOffset`,
  `ListLimit` — compose list queries without leaking storage primitives.

## Rules

- Imports `api/...` only for error reason enums. Never imports `service`,
  never imports `data`.
- `biz.New` (see `biz.go`) builds all usecases from their repo deps;
  `buildApp` in `cmd/<app>/main.go` calls it.

See `AGENTS.md` for the full layering contract.