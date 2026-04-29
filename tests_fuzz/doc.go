// Package fuzz contains fuzz tests for the ocpp16messages module.
//
// # What It Means
//
// This package exercises every public constructor in the module —
// Req(), Conf(), and New*() functions — with machine-generated inputs
// to discover inputs that panic, violate error-sentinel contracts, or
// break success invariants such as round-tripping, UTC-only DateTime
// values, and nil-vs-empty slice preservation.
// Enum fuzzers verify that IsValid() returns true if and only if the input
// matches an OCPP-spec value.
//
// # When It Is Used
//
// Fuzz tests are opt-in and are not executed during a normal go test ./... run.
// They run on a weekly schedule in CI and can be invoked locally via:
//
//	make test-fuzz
//
// To run a single fuzzer directly:
//
//	go test -tags=fuzz -fuzz=^FuzzAuthorizeReq$ -fuzztime=10s ./tests_fuzz
//
// Tune the time budget with FUZZTIME (default 5s) and parallelism with
// FUZZPROCS (default 4).
//
// # What It Is Not
//
// This package is not part of the public API and is not imported by application
// code. It does not run under go test ./... without the fuzz build tag. It is
// not a replacement for unit tests in the per-package tests/ subdirectories;
// unit tests cover specific known-good and known-bad inputs while fuzz tests
// explore the unknown space around them.
//
// # Adjacent Concepts
//
//   - tests_race: the opt-in race-detector suite that verifies immutability and
//     aliasing contracts; complementary to fuzz coverage.
//   - tests_benchmark: the opt-in benchmark suite for performance regression
//     detection.
//   - Each message package's tests/ subdirectory: the standard unit tests that
//     run on every CI push.
package fuzz
