package errorhandling

import (
	"fmt"

	exercise "github.com/nathancastelein/go-test-production/exercise/errorhandling"
)

func FizzbuzzWithError(input int) (string, error) {
	if input == 0 {
		return "", fmt.Errorf("input cannot be equal to 0: %d", input)
	}

	if input < 0 {
		return "", fmt.Errorf("input cannot be negative: %d", input)
	}

	return exercise.Fizzbuzz(input), nil
}
