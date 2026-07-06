# Service

Transport adapter layer: converts DTO ↔ DO at the boundary and delegates
to `biz` usecases. One file per resource.

- `convert<Resource>(req)` parses proto → DO; build the reply inline at the
  return site (`&v1.<Resource>{...}`, `*v1.<Resources>Set`, or
  `&emptypb.Empty{}` for deletes) so each handler stays self-contained.
- Embed `v1.Unimplemented<Resource>ServiceServer` on the service struct.
- Parse AIP list requests via `filtering` / `ordering` / `pagination`; apply
  `fieldmask.Update` for partial updates (`UpdateTodo` in `todo.go`).
- Validate request inputs at the boundary before delegating to the usecase.
- Return `biz` errors. No business rules, no storage access, no PO.
- Imports: `api/...` (DTO) and `biz` (DO) only — never `data`.
- `service.New` (see `service.go`) builds all services from their usecase
  deps; called by `buildApp` in `cmd/<app>/main.go`.

See `AGENTS.md` for the full layering contract.