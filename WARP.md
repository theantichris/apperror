# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

**apperror** is a Go error handling library that provides structured error management with:

- **AppError struct**: Core error type with Type, Message, and Cause fields
- **Error wrapping**: Chains errors while preserving the original cause
- **Double-wrapping prevention**: Prevents the same error type from being wrapped multiple times
- **Standard error interface**: Implements Go's standard error interface

## Development Commands

### Initial Setup

```bash
# Download dependencies
go mod download

# Tidy up dependencies
go mod tidy
```

### Building and Testing

```bash
# Build the library
go build ./...

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run specific test
go test -run TestSpecificFunction ./...

# Run benchmarks
go test -bench=. ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run golint (if installed)
golint ./...

# Run staticcheck (if installed)
staticcheck ./...
```

## Architecture

### Core Components

**AppError Struct**:

```go
type AppError struct {
    Type    error  // The error type
    Message string // Descriptive message
    Cause   error  // Underlying error, if any
}
```

**Key Features**:

- **Error Wrapping**: The `New()` function creates structured errors with optional cause chaining
- **Double-wrapping Prevention**: Automatically prevents wrapping the same error type multiple times
- **Error Interface**: Implements standard Go error interface with intelligent message formatting
- **Cause Chaining**: When a cause is present, error messages are formatted as "message: cause"

**Usage Pattern**:

```go
// Create a new structured error
baseErr := errors.New("database connection failed")
appErr := apperror.New(baseErr, "failed to save user", baseErr)

// Error message will be: "failed to save user: database connection failed"
fmt.Println(appErr.Error())
```

## File Organization

Current structure:

- `apperror.go`: Main library implementation with AppError struct and New() function
- `apperror_test.go`: Comprehensive tests covering error creation, wrapping, and formatting
- `go.mod`: Module definition (`github.com/theantichris/apperror`)
- `.vscode/settings.json`: VS Code configuration (includes "apperror" in spell check dictionary)

## Testing Strategy

Current test coverage includes:

- **Error Creation**: Tests for `New()` function with various combinations of type, message, and cause
- **Double-wrapping Prevention**: Verifies that wrapping the same error type returns the original error
- **Error Formatting**: Tests both simple messages and cause-chained messages
- **Parallel Testing**: All tests are designed to run in parallel for better performance

**Test Structure**:

- Tests use table-driven patterns with subtests
- Each major function has dedicated test functions
- Tests verify both successful operations and edge cases
