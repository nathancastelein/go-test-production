package errorhandling

import (
	"fmt"
	"os"
)

func WriteFile() error {
	file, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		return err
	}

	for idx := range 100 {
		res, err := FizzbuzzWithError(idx)
		if err != nil {
			return err
		}

		fmt.Fprintf(file, "%d: %s", idx, res)
	}

	return file.Close()
}
