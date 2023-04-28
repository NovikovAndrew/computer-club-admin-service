package configuration

import (
	"context"

	"github.com/NovikovAndrew/computer-club-admin-service.git/internal/adapters/api/componenthandler"
)

type service struct {
	storage          Storage
	componentService componenthandler.ComponentService
}

func NewConfigurationService(storage Storage, componentService componenthandler.ComponentService) *service {
	return &service{
		storage:          storage,
		componentService: componentService,
	}
}

func (s *service) GetOne(ctx context.Context, id string) (*ComputerConfiguration, error) {
	panic("impl")
}

func (s *service) GetAllByComputerID(ctx context.Context, id string) ([]ComputerConfiguration, error) {
	panic("impl")
}

func (s *service) Create(ctx context.Context, cfg CreateComputerConfigurationDTO) error {
	panic("impl")
}

func (s *service) Update(ctx context.Context, cfg UpdateComputerConfigurationDTO) error {
	panic("impl")
}

func (s *service) Delete(ctx context.Context, id string) error {
	panic("impl")
}
