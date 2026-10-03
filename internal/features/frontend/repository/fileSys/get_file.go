package frontend_repository_filesys

import (
	"errors"
	"fmt"
	"os"

	core_errors "github.com/qury53/finapp/internal/core/transport/http/errors"
)

func (r *FrontendRepository) GetFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf(
				"repository: failed to read file with name: %s, it doesn't exist: %w",
				filename,
				core_errors.ErrNotFound,
			)
		}

		return nil, fmt.Errorf(
			"repository: failed to read file: %w",
			err,
		)
	}

	return file, nil
}
