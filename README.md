# Kratos Service Template

A template for creating new Kratos services with HTTP and gRPC
transports, protobuf-first APIs, manual layer composition, and OpenAPI
plus proto doc generation. A small CRUD example (Todo resource) shows
the API shape, layering, code generation, and testing conventions —
replace it with your own domain model when creating a real project.

## Create a New Project

1. Copy or generate a repository from this template.
2. Run the interactive setup; it prompts for an app name and rewrites
   the Go module path plus all import references:

   ```bash
   mise run setup
   ```

3. Replace the sample CRUD resource with your own.
4. Regenerate code and verify the build:

   ```bash
   mise run generate
   go test ./...
   ```

## What Is Included

- Kratos HTTP and gRPC server setup.
- Protobuf API definitions with generated Go stubs.
- OpenAPI document and Markdown proto docs, generated into `docs/`.
- Generated mocks for service and repo tests, in `mock/`.
- Manual dependency composition via `buildApp` in `cmd/<app>/main.go`.
- Layered `service`, `biz`, and `data` packages.
- An in-memory repository for the sample resource.
- Server-streaming (`WatchTodos`) and bidirectional (`SyncTodos`) RPCs.
- `mise`-driven tasks for generation, lint, format, and builds.

## Project Layout

```text
api/                  Protobuf APIs and generated bindings
cmd/                  Application entrypoints
configs/              Local configuration
docs/                 Generated OpenAPI document and proto docs
mock/                 Generated mocks for service/repo tests
internal/conf/        Config protobufs
internal/server/      HTTP and gRPC server construction
internal/service/     Transport adapters; DTO ↔ DO conversion
internal/biz/         Usecases, domain models, errors, repo interfaces
internal/data/        Repository implementations and storage clients
mise.toml             mise tool versions and tasks
Dockerfile            Multi-stage, cross-compiling image
```

## API Template Practices

The sample Todo API demonstrates the conventions this template expects:

- Resource-oriented CRUD: create, get, list, update, delete.
- HTTP annotations via `google.api.http`.
- Required fields marked with `google.api.field_behavior`.
- List requests with `page_size`, `page_token`, `filter`, `order_by`.
- Pagination via `go.einride.tech/aip/pagination`.
- Partial updates with `google.protobuf.FieldMask` and `fieldmask.Update`.
- Streaming RPC definitions for server and bidirectional streams.

The in-memory data layer is intentionally simple — it shows flow across
layers, not a full query engine. Real repositories apply parsed filters
and ordering in SQL, Ent, or another storage layer.

## Development Commands

Tooling and generators are managed by `mise` (see `mise.toml`).

Install declared tools:

```bash
mise install
```

Regenerate proto stubs, mocks, OpenAPI, and proto docs:

```bash
mise run generate
```

Build the binary (goreleaser snapshot):

```bash
mise run build:binary
```

Build and push the multi-arch Docker image (uses `REGISTRY`, `PKG_PATH`,
`APP_NAME` envs):

```bash
mise run build:docker
```

Lint and format:

```bash
mise run lint
mise run fmt
```

Test:

```bash
go test ./...
```

## Run Locally

```bash
go run ./cmd/server -conf ./configs
```

Default local ports (configured in `configs/config.yaml`):

- HTTP: `0.0.0.0:8000`
- gRPC: `0.0.0.0:9000`

## Docker

Local single-arch build and run:

```bash
docker build -t <your-image-name> .
docker run --rm -p 8000:8000 -p 9000:9000 \
  -v </path/to/your/configs>:/data/conf \
  <your-image-name>
```

Multi-arch build and push:

```bash
mise run build:docker
```
