package frontend_service

type FrontendService struct {
	FrontendRepository FrontendRepository
}

type FrontendRepository interface {
	GetFile(filename string) ([]byte, error)
}

func NewFrontendService(repository FrontendRepository) *FrontendService {
	return &FrontendService{
		FrontendRepository: repository,
	}
}
