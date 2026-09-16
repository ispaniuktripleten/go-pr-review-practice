# Go PR review fixtures
This is a teaching repository with five independent review branches, each based on main.
The base contains model, API adapter, and repository packages. No application deployment
or credentials are required. Use a supported Go release (language baseline in go.mod: 1.22).
Validate locally using go test ./... and go vet ./... .
PR-05 contains an intentionally non-isolated test skipped by default. Keep RUN_LIVE_TESTS
unset. Credential strings in review branches are synthetic examples, never real secrets.
