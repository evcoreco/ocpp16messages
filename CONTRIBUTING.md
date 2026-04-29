# Contributing

Thanks for your interest in improving `ocpp16messages`.

## What It Means

A contribution to this project is any change — code, documentation, tests, or
bug reports — that improves correctness, clarity, or coverage of the OCPP 1.6
message set. Every public type, constructor, and piece of documentation is part
of the library's contract with its users, so changes are held to the same
standard as the existing code.

## When It Is Used

Open a pull request when you have identified a bug, a missing behavior from the
OCPP 1.6 specification, a documentation gap, or a correctness issue in an
existing test. Before submitting, run:

```sh
make format
make lint
make test
```

### OCPP compliance changes

When changing validation or message fields, include:

- The OCPP 1.6 operation name and whether it affects `.req` and/or `.conf`.
- The spec reference (section and page number) or a short excerpt.
- A minimal failing test demonstrating the previous behavior.
- Updated tests proving the new behavior and no regressions.

### Adding new public API

- Add public API tests in `tests/` subdirectories using `package <name>_test`.
- Add examples only when they improve discoverability.
- Follow the constructor and validation pattern (`Req`, `Conf`, `New*`).
- Document every exported symbol following the four-section pattern:
  what it means, when it is used, what it is not, adjacent concepts.

### Concurrency and immutability

If your change touches pointers or slices:

- Ensure constructors do not store caller-owned references.
- Update `./tests_race` to cover the new or changed aliasing behavior.

## What It Is Not

A contribution is not a place to add features outside the OCPP 1.6 scope, to
introduce transport or WebSocket logic, or to re-implement standard operations
via DataTransfer. It is not acceptable to skip the opt-in test suites for
changes that touch constructors, pointers, or slices — those suites exist
precisely to guard the contracts that are hardest to verify in unit tests.

## Adjacent Concepts

- `ADDING_MESSAGE.md` — step-by-step guide for adding a new OCPP operation.
- `CLAUDE.md` — detailed development guidelines for this repository.
- `tests_fuzz/` — fuzz suite; update when adding new constructors.
- `tests_race/` — race-detector suite; update when adding pointer or slice
  fields.
- Opt-in suites: `make test-fuzz`, `make test-race`, `make test-bench`.

## Pull Requests

- Keep PRs focused and small where possible.
- Ensure Markdown changes pass `markdownlint`.
- Prefer atomic tests: one behavior per test function.

## Code of Conduct

This project follows `CODE_OF_CONDUCT.md`.
