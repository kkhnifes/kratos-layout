# Data

Storage layer: implements `biz.<Resource>Repo` interfaces, owns PO shapes
and long-lived storage clients. Imports `biz`; never imports `service` or
DTOs (`api/...`). See `AGENTS.md` for the full layering contract.

- **`*Data`** (`data.go`): shared struct holding storage clients (db, redis,
  …). Constructed once by `NewData` and passed to every repo; repos never
  build their own clients.
- **Repo constructor**: `func New<Resource>Repo(d *Data) biz.<Resource>Repo`
  — returns the interface, never the concrete type. Wired in `buildApp`
  (`cmd/<app>/main.go`).
- **PO**: define a Persistent Object only when storage shape diverges from
  the DO. PO types stay in `data`; convert via free functions
  `new<Resource>(do)` (write, DO→PO) and `toBiz(po)` (read, PO→DO). When
  storage shape matches the DO, skip the PO and use the DO directly (see
  `todo.go`, an in-memory repo with no PO).
- **Querying**: translate `biz.ListOptions` (`Filter`, `OrderBy`, `Offset`,
  `Limit`) into the storage driver's query language inside the repo.
- **Errors**: map driver errors to `biz` typed errors
  (`biz.Err<Resource><Cause>`) so callers above never branch on the driver.