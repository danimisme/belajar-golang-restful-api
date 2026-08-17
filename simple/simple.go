package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed to create service")
	} else {
		return &SimpleService{SimpleRepository: repository}, nil
	}
}