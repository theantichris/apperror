package apperror

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("creates a new AppErr", func(t *testing.T) {
		t.Parallel()

		err := errors.New("resource not found")

		want := &AppError{
			Type:  err,
			Cause: nil,
		}

		got := New(err, "resource not found", nil)

		if got.Type != want.Type {
			t.Errorf("got %v, want %v", got.Type, want.Type)
		}

		if got.Cause != want.Cause {
			t.Errorf("got %v, want %v", got.Cause, want.Cause)
		}
	})

	t.Run("prevents double-wrapping of the same error type", func(t *testing.T) {
		t.Parallel()

		baseErr := errors.New("database error")

		want := New(baseErr, "failed to connect to database", nil)
		got := New(baseErr, "another database error occurred", want)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestError(t *testing.T) {
	t.Run("returns message", func(t *testing.T) {
		t.Parallel()

		got := &AppError{
			Type:    errors.New("resource not found"),
			Message: "resource not found",
			Cause:   nil,
		}

		want := "resource not found"

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns message when no cause", func(t *testing.T) {
		t.Parallel()

		cause := errors.New("root cause")
		appErr := &AppError{
			Type:    cause,
			Message: "something failed",
			Cause:   cause,
		}

		got := appErr.Error()
		want := "something failed: root cause"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
