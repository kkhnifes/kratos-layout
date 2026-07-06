# conf

Config protobufs for the service template.

- `conf.proto` — source. Defines `Bootstrap` (Server, Data, Log) loaded by `cmd/<app>/main.go` to parse `configs/config.yaml`.
- `conf.pb.go`, `conf.pb.validate.go` — generated. Never hand-edit.
- `buf.yaml`, `buf.lock`, `buf.gen.yaml` — buf config for this module.

## Regenerate

```sh
mise run generate
```

Runs `buf generate` at the repo root, covering both `api/` and `internal/conf/` (see `mise.toml` `generate:buf`).