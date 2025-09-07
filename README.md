# apperror

A structured error handling library for Go that provides enhanced error management with type safety, cause chaining, and intelligent formatting.

## Features

- 🏗️ **Structured Errors**: Define errors with type, message, and optional cause
- 🔗 **Error Chaining**: Chain errors while preserving the original cause
- 🛡️ **Double-wrapping Prevention**: Automatically prevents wrapping the same error type multiple times
- 📝 **Intelligent Formatting**: Smart error message formatting with cause integration
- ✅ **Standard Interface**: Implements Go's standard `error` interface
- 🚀 **Zero Dependencies**: No external dependencies, pure Go standard library

## Installation

```bash
go get github.com/theantichris/apperror
```

## Quick Start

```go
package main

import (
    "errors"
    "fmt"
    "github.com/theantichris/apperror"
)

func main() {
    // Create a base error type
    dbError := errors.New("database error")

    // Create a structured error with cause
    connectionErr := errors.New("connection timeout")
    appErr := apperror.New(dbError, "failed to connect to database", connectionErr)

    fmt.Println(appErr.Error())
    // Output: failed to connect to database: connection timeout
}
```

## Usage

### Creating Errors

```go
// Simple error without cause
validationErr := errors.New("validation error")
appErr := apperror.New(validationErr, "invalid user input", nil)
fmt.Println(appErr.Error()) // Output: invalid user input

// Error with cause
networkErr := errors.New("network error")
timeoutErr := errors.New("request timeout")
appErr := apperror.New(networkErr, "API call failed", timeoutErr)
fmt.Println(appErr.Error()) // Output: API call failed: request timeout
```

### Double-wrapping Prevention

The library automatically prevents double-wrapping of the same error type:

```go
baseErr := errors.New("database error")
firstWrap := apperror.New(baseErr, "first wrap", nil)
secondWrap := apperror.New(baseErr, "second wrap", firstWrap)

// secondWrap will be the same as firstWrap
fmt.Println(firstWrap == secondWrap) // Output: true
```

### Error Structure

The `AppError` struct provides three fields:

```go
type AppError struct {
    Type    error  // The error type/category
    Message string // Descriptive message
    Cause   error  // Underlying error, if any
}
```

### Accessing Error Information

```go
baseErr := errors.New("file error")
appErr := apperror.New(baseErr, "failed to read config", errors.New("file not found"))

// Access the error type
fmt.Println(appErr.Type.Error()) // Output: file error

// Access the message
fmt.Println(appErr.Message) // Output: failed to read config

// Access the cause
if appErr.Cause != nil {
    fmt.Println(appErr.Cause.Error()) // Output: file not found
}
```

## API Reference

### Types

#### `AppError`

```go
type AppError struct {
    Type    error  // The error type
    Message string // Descriptive message
    Cause   error  // Underlying error, if any
}
```

### Functions

#### `New(errorType error, message string, cause error) *AppError`

Creates a new `AppError` with the specified type, message, and optional cause.

**Parameters:**

- `errorType`: The error type/category
- `message`: Descriptive error message
- `cause`: Optional underlying error (can be `nil`)

**Returns:** `*AppError`

**Behavior:**

- If `cause` is already an `AppError` with the same `Type`, returns the existing `AppError` to prevent double-wrapping
- Otherwise, creates a new `AppError` instance

### Methods

#### `(err *AppError) Error() string`

Implements the standard Go `error` interface.

**Returns:**

- If `Cause` is `nil`: returns `Message`
- If `Cause` is not `nil`: returns `"Message: Cause.Error()"`

## Testing

Run the tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
