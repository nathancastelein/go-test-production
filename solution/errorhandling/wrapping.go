package errorhandling

import (
	"errors"
	"fmt"

	exercise "github.com/nathancastelein/go-test-production/exercise/errorhandling"
)

var ErrInvalidInput = errors.New("invalid input")

func FizzbuzzWithWrappingError(input int) (string, error) {
	if input == 0 {
		return "", fmt.Errorf("%w: input equals 0", ErrInvalidInput)
	}

	if input < 0 {
		return "", fmt.Errorf("%w: input is negative: %d", ErrInvalidInput, input)
	}

	return exercise.Fizzbuzz(input), nil
}
