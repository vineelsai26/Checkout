# Checkout

Checkout is a small Go CLI for moving personal projects between the source tree
and a local checkout workspace. It is intentionally agent-friendly: commands
return normal errors instead of panicking for expected CLI problems, dry runs are
available, and local dependency/cache directories are skipped by default.

## Usage

```sh
checkout init [flags] [project]
checkout deinit [flags] <project> [source-folder]
```

`init` moves a project from the configured source tree into `~/Personal`.
`deinit` moves it back under the source tree. Both commands refuse to overwrite
an existing destination.

Useful flags:

```sh
checkout init --dry-run my-app
checkout init --include-env my-app
checkout init --exclude tmp --exclude '*.log' my-app
checkout init --no-open my-app
checkout init --source-dir /tmp/source --checkout-root /tmp/checkout my-app
```

By default Checkout excludes local dependency and cache folders such as
`node_modules`, `.next`, `dist`, `build`, `.turbo`, `.venv`, `venv`,
`__pycache__`, `.pytest_cache`, `.mypy_cache`, `.ruff_cache`, `.tox`, `.nox`,
and `target`.

Secret-bearing `.env` and `.env.*` files are excluded by default. Template files
such as `.env.example`, `.env.template`, and `.env.sample` are still copied.
Use `--include-env` when the environment files are intentionally part of the
move.

After a successful copy, Checkout deletes only source entries that were copied.
Excluded `.env` files and local cache directories remain in the original source
tree, so an ordinary move cannot silently discard omitted local state. Copy or
cleanup errors abort the operation and leave a resumable source tree.

The default source root comes from `CHECKOUT_SOURCE_DIR`, then
`~/.checkout/source_dir`, then `/Users/vineel/Dropbox/GitHub`. The default
checkout root comes from `CHECKOUT_ROOT`, then `~/Personal`. Both can be
overridden per command with `--source-dir` and `--checkout-root`.

## Provenance

- Imported from `https://github.com/vineelsai26/Checkout`
- Imported commit: `6663036`
- Monorepo path: `tools/checkout`

## Commands

```sh
go test ./...
go build
```

The original repository did not include a README. This note records the import
origin and the local verification commands used in `vstack`.
