package frontend_service

import (
	"fmt"
	"os"
	"path"
)

func (s *FrontendService) GetProfilePage() ([]byte, error) {
	filePath := path.Join(os.Getenv("PROJECT_ROOT"), "frontend/profile.html")

	file, err := s.FrontendRepository.GetFile(filePath)
	if err != nil {
		return nil, fmt.Errorf(
			"service: failed to get profile page: %w",
			err,
		)
	}

	return file, nil
}
