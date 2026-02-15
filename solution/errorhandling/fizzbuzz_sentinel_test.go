package errorhandling

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyFizzbuzzProcess(t *testing.T) {
	// Arrange
	testCases := []struct {
		name        string
		input       int
		expectError bool
	}{
		{name: "should work with positive number", input: 1, expectError: false},
		{name: "should work with zero", input: 0, expectError: false},
		{name: "should fail with negative number", input: -1, expectError: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			err := MyFizzbuzzProcess(tc.input)

			// Assert
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
