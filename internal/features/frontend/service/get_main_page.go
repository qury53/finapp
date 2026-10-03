package frontend_service

import (
	"fmt"
	"os"
	"path"
)

func (s *FrontendService) GetMainPage() ([]byte, error) {
	filePath := path.Join(os.Getenv("PROJECT_ROOT"), "frontend/index.html")
	file, err := s.FrontendRepository.GetFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(
			"service: failed to get file: %w",
			err,
		)
	}

	return file, nil
}
