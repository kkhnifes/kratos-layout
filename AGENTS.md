# Repository Guidelines

Kratos service template. This file is the layering contract agents must follow.

## Project structure

```
api/<domain>/<version>/   Proto sources and generated stubs. Public contract.
cmd/<app>/                Entrypoint; `buildApp` in main.go composes all layers.
configs/                  Runtime config (config.yaml). No secrets.
internal/conf/            Config proto and generated stubs.
internal/pkg/             Shared helpers (e.g. logging).
internal/server/          HTTP/gRPC server wiring.
internal/service/         Transport adapters; one file per resource.
internal/biz/             Domain models, usecases, repo interfaces, errors.
internal/data/            Repo implementations and storage clients.
mock/                     Generated mocks (mockery). Never hand-edit.
docs/                     Generated OpenAPI + proto docs.
```

## Build & tooling

`mise` manages tools and tasks (see `mise.toml`):

- `mise run generate` — `buf generate` + `mockery` (proto stubs, mocks, OpenAPI, docs).
- `mise run lint` / `mise run fmt` — golangci-lint (+ yamllint) / `golangci-lint fmt`.
- `mise run build:binary` — goreleaser snapshot. `mise run build:docker` — multi-arch push.
- `mise run setup` — interactive: prompts app name, rewrites module path and imports.

No Makefile, no Wire. Dependency composition is manual in `buildApp`.

## Layering & dependency rules

Three model shapes flow through three layers. `biz` owns the DO, `data` owns
the PO; `service` is a pass-through that converts at its boundary.

```
   client ──► DTO ──► service ──► DO ──► biz ──► DO ──► data ──► PO ──► storage
                                  ▲                ▲
                                  │ declares       │ implements
                                  └─── repo IF ────┘

   DTO  Data Transfer Object — proto request / response.
   DO   Domain Object        — pure biz model, no proto, no storage tags.
   PO   Persistent Object    — storage shape, owned by `data`.
```

| Layer   | Owns | Speaks at boundary | Never speaks            |
|---------|------|--------------------|-------------------------|
| service | —    | DTO ↔ DO          | PO, storage client      |
| biz     | DO   | DO                 | DTO, PO, storage client  |
| data    | PO   | DO ↔ PO           | DTO                     |

- `service` imports `api/...` (DTO) and `biz` (DO). Never `data`.
- `biz` imports `api/...` only for error reason enums. Never `service`,
  never `data`. The repo interface declared here is the inversion seam.
- `data` imports `biz` to implement the repo interface. Never `service`,
  never DTOs.
- `cmd` is the only place that composes all layers. Each layer exposes a
  `New` initializer (`biz.New`, `service.New`, `server.New`) that
  `buildApp` calls in dependency order: `data.NewData` →
  `data.New<Resource>Repo` → `biz.New` → `service.New` → `server.New`.

A change crossing these arrows the wrong way is a layering bug; fix the
design rather than add the import.

### Layer responsibilities

**service (DTO ↔ DO)**

- `convert<Resource>` parses an incoming proto into a DO. The reverse
  direction is built inline at the return site — the reply is whatever
  the proto declares (usually `&v1.<Resource>{...}`, sometimes
  `*v1.<Resources>Set`, or `&emptypb.Empty{}` for deletes). Inlining
  keeps each handler self-contained.
- Embed `Unimplemented<Resource>ServiceServer`.
- Parse AIP list requests via `filtering` / `ordering` / `pagination`;
  apply `fieldmask.Update` for partial updates.
- Validate request inputs at the boundary before delegating to the
  usecase. No business rules, no storage access, no PO.

**biz (DO only)**

- Owns the DO (`type <Resource> struct` — no proto, no storage tags),
  the usecase, and the repo interface (`type <Resource>Repo interface`).
- Owns typed errors built with `errors.NotFound` / `errors.BadRequest`
  plus the API error reason enum.
- Owns `ListOption` helpers — `ListFilter`, `ListOrderBy`, `ListOffset`,
  `ListLimit` — so callers compose queries without leaking storage
  primitives.

**data (DO ↔ PO)**

- Implement `biz.<Resource>Repo`. The constructor returns the interface:
  `func New<Resource>Repo(d *Data) biz.<Resource>Repo`.
- Define a PO when storage shape diverges from the DO. PO types stay
  inside `data`. Convert with free functions `new<Resource>` (DO → PO,
  write) and `toBiz` (PO → DO, read). Driver-specific builders never
  leave `data`.
- `*Data` (in `internal/data/data.go`) holds long-lived storage clients.
  Repos receive `*Data` and never construct their own clients.
- Translate `ListOptions.Filter` / `ListOptions.OrderBy` into the
  driver's query language inside the repo. Map driver errors to `biz`
  typed errors so callers above never branch on the driver.

**server**

- Construct HTTP/gRPC servers, apply middleware, register services. No
  translation, no business logic.

### Add-a-resource checklist

1. **DTO**: define `Create<Resource>` / `Get<Resource>` /
   `List<Resources>` / `Update<Resource>` / `Delete<Resource>` in
   `api/<domain>/<version>/`, then `mise run generate`.
2. **DO + repo interface**: declare both in `biz`; build the usecase on
   top of the interface.
3. **Repo impl**: implement in `data` returning `biz.<Resource>Repo`;
   add a PO and conversion helpers when storage shape diverges from DO.
4. **Wiring**: add the repo constructor to `buildApp`
   (`cmd/<app>/main.go`); add the usecase to `biz.New`
   (`internal/biz/biz.go`); add the service to `service.New`
   (`internal/service/service.go`); register HTTP/gRPC services in
   `server.New` (`internal/server/server.go`).
5. **Regenerate**: `mise run generate` to refresh proto stubs and mocks.

### Testing seam

Tests live beside the code they cover (`*_test.go`). Test layers in
isolation: service tests fake the usecase, biz tests fake the repo, data
tests exercise repo implementations at the storage boundary. Mocks are
generated by mockery into `mock/` (package `<pkg>mock`).

## Generation & generated files

Regenerate via `mise run generate`. Never hand-edit `*.pb.go`,
`*_grpc.pb.go`, `*_http.pb.go`, `*.pb.validate.go`, or files under
`mock/`. Regenerated files belong in the same commit as their source.

## Naming & error reasons

- Resource: `<Resource>` (e.g., `Todo`); collection RPC: `List<Resources>`.
- Types: repo `<Resource>Repo`, usecase `<Resource>Usecase`, service
  `<Resource>Service`. PO types live inside `internal/data/`; convert
  with `new<Resource>(do)` / `toBiz(po)` free functions.
- Error reasons: declared in `api/<domain>/<version>/error_reason.proto`,
  surfaced as `Err<Resource><Cause>` in `biz`.

## Commits & security

- Conventional Commits: `feat:`, `fix:`, `refactor:`, `chore(deps):`,
  `docs:`, `test:`.
- Never commit real credentials in `configs/config.yaml`.