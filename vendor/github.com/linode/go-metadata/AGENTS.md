# Agent Guidelines

## Project Shape
- This repository is a single-package Go module, `github.com/linode/go-metadata`, for the Linode Metadata Service client.
- Keep public API work in the root `metadata` package; examples live under [examples/](examples/) and integration coverage lives in the separate module under [test/integration/](test/integration/).
- Link to existing docs instead of copying them: start with [README.md](README.md) for setup and usage, and [CONTRIBUTING.md](CONTRIBUTING.md) for contribution expectations.

## Commands Agents Should Use
- Unit tests: `make unit-test` or `go test -v ./`
- Vet: `go vet ./...`
- Lint: `make lint` (`SKIP_DOCKER=1 make lint` runs local `golangci-lint`; default runs the Docker image)
- Format: `make fmt` (`gofumpt -w -l .`) or `make fix-lint` for formatter plus linter fixes
- Dependency hygiene: `go mod tidy` followed by checking for a clean diff
- Remote E2E: `make test-deps` then `make e2e` with `LINODE_TOKEN` and an SSH public key available; this provisions paid Linode infrastructure unless cleanup is enabled.
- Local E2E: `make e2e-local`, but only from within a Linode instance. It delegates to [test/integration/Makefile](test/integration/Makefile) and writes JUnit XML reports.

## Code Conventions
- Follow the existing functional-option pattern: `ClientOption`, `TokenOption`, and `WatcherOption` mutate private config structs.
- Public client methods should accept `context.Context`; tests and examples use `context.Background()` at call sites.
- Endpoint files are organized by resource: [instance.go](instance.go), [network.go](network.go), [token.go](token.go), [userdata.go](userdata.go), and [sshkeys.go](sshkeys.go).
- Watchers are thin resource-specific facades over the generic polling implementation in [watcher_generic.go](watcher_generic.go). Preserve channel semantics: `Updates` and `Errors` are exposed channels, `Start` blocks until cancellation, and `Close` signals the poller.
- The HTTP layer uses `resty`; keep client setup and token refresh behavior centralized in [client.go](client.go).
- Use the Go version and toolchain declared in [go.mod](go.mod) for development tasks.

## Testing Notes
- Unit tests are co-located with root package files and should not require live Linode infrastructure.
- Integration tests use their own module in [test/integration/go.mod](test/integration/go.mod) with a local replace back to this repo.
- Prefer `httpmock`/test helpers for API behavior in tests unless the change specifically targets live metadata behavior.
- Be careful with E2E cleanup: `CLEANUP_TEST_LINODE_INSTANCE=true` removes provisioned infrastructure; leaving it false can incur ongoing costs.

## CI And Review Gotchas
- CI runs lint, `go vet ./...`, `go mod tidy`, then fails if the working tree changes.
- PR titles are checked in [.github/workflows/ci.yml](.github/workflows/ci.yml); non-exempt PRs need a `TPT-1234:`-style prefix.
- The linter config in [.golangci.yml](.golangci.yml) enables `gosec` and formatters `gofumpt`/`goimports`, with examples excluded from formatter and linter paths.
