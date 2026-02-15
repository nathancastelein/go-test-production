# Go Testing + Production - Exercises

This repository contains the exercises for Module 2 of the Go Training course.

## Prerequisites

- Go 1.25 or later
- Docker (for integration tests with testcontainers)

## Repository Structure

```
go-test-production/
├── exercise/           # Exercise stubs - implement these!
│   ├── unit/fizzbuzz/  # Unit testing exercises
│   ├── mocking/http/   # HTTP mocking with gock
│   ├── mocking/sql/    # SQL mocking with sqlmock
│   ├── errorhandling/  # Error handling patterns
│   └── logging/        # Structured logging with slog
├── solution/           # Complete solutions
├── examples/           # Demo code and working examples
│   └── integration/sql/ # Testcontainers integration example
```

## Exercise Order

### Unit Testing (FizzBuzz)

| Exercise           | File                                          | Description                      |
| ------------------ | --------------------------------------------- | -------------------------------- |
| Basic Test         | `exercise/unit/fizzbuzz/fizzbuzz_test.go`     | Write your first unit test       |
| Table-Driven Test  | `exercise/unit/fizzbuzz/fizzbuzz_tdt_test.go` | Use table-driven test pattern    |
| TDT with Subtests  | `exercise/unit/fizzbuzz/fizzbuzz_tdt_run_test.go` | Use t.Run for subtests      |
| Fuzz Test          | `exercise/unit/fizzbuzz/fizzbuzz_fuzzing_test.go` | Write fuzz tests             |

### Mocking (HTTP)

| Exercise      | File                              | Description                    |
| ------------- | --------------------------------- | ------------------------------ |
| Mock HTTP API | `exercise/mocking/http/http.go`   | Mock external HTTP calls       |

### Mocking (SQL)

| Exercise         | File                                      | Description                   |
| ---------------- | ----------------------------------------- | ----------------------------- |
| Mock SQL Queries | `exercise/mocking/sql/sql.go`             | Mock database queries         |
| SQL with Testify | `exercise/mocking/sql/sql_testsuite_test.go` | Use testify suite          |

### Error Handling

| Exercise        | File                                          | Description                      |
| --------------- | --------------------------------------------- | -------------------------------- |
| Basic Errors    | `exercise/errorhandling/fizzbuzz.go`          | Create and return errors         |
| Sentinel Errors | `exercise/errorhandling/fizzbuzz_sentinel.go` | Use sentinel error pattern       |
| Custom Errors   | `exercise/errorhandling/fizzbuzz_custom_error.go` | Implement error interface    |
| Defer & Panic   | `exercise/errorhandling/defer.go`             | Use defer for cleanup            |

### Logging (slog)

| Exercise           | File                             | Description                       |
| ------------------ | -------------------------------- | --------------------------------- |
| Logging Middleware | `exercise/logging/middleware.go` | Add structured logging            |
| Context Logger     | `exercise/logging/context.go`    | Propagate logger via context      |
| Benchmarks         | `exercise/logging/middleware_benchmark_test.go` | Benchmark logging      |

See `exercise/logging/README.md` for detailed exercise instructions.

## Running Tests

### Test your exercise implementation

```bash
# Test a specific exercise
go test ./exercise/unit/fizzbuzz/... -run ^TestFizzbuzz$

# Test all exercises (will fail until implemented)
go test ./exercise/...
```

### Verify against solutions

```bash
# All solution tests should pass
go test ./solution/...
```

### Run integration tests (requires Docker)

```bash
go test -tags=integration ./examples/integration/...
```

## Go Version

This module requires Go 1.25 or later. Check your version:

```bash
go version
```
