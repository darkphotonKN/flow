package fileService

import (
	"fmt"

	"github.com/google/uuid"
)

type ApiData struct {
	StatusCode int
	Message    string
	Result     []string
}

func example() {
	service := NewService()

	fileService := NewFileService("asdasd", service)
}

// SERVICE

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (p *Service) create(test int) error {
	fmt.Println("Create")
	if test > 1 {
		return nil
	}
	return fmt.Errorf("Error happened!")
}

// FILE SERVICE

type FileService struct {
	ID      uuid.UUID
	Name    string
	Service *Service
}

func NewFileService(name string, service *Service) *FileService {
	return &FileService{
		ID:      uuid.New(),
		Name:    name,
		Service: service,
	}
}

func (fs *FileService) downloadFile() {
	// download file

	err := fs.Service.create(2)
}
