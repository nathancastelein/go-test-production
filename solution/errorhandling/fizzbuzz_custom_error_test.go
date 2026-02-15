package errorhandling

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzzWithCustomError(t *testing.T) {
	// Arrange
	testCases := []struct {
		name        string
		input       int
		expectError bool
	}{
		{name: "should work with positive number", input: 1, expectError: false},
		{name: "should fail with zero", input: 0, expectError: true},
		{name: "should fail with negative number", input: -3, expectError: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			_, err := FizzbuzzWithCustomError(tc.input)

			// Assert
			if tc.expectError {
				require.Error(t, err)

				var invalidInput InvalidInput
				require.True(t, errors.As(err, &invalidInput), "expected error type InvalidInput, got %T", err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
