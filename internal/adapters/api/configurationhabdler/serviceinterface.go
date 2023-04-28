package configurationhabdler

import (
	"context"

	"github.com/NovikovAndrew/computer-club-admin-service.git/internal/domain/configuration"
)

type ConfigurationService interface {
	GetOne(ctx context.Context, id string) (*configuration.ComputerConfiguration, error)
	GetAllByComputerID(ctx context.Context, id string) ([]configuration.ComputerConfiguration, error)
	Create(ctx context.Context, cfg configuration.CreateComputerConfigurationDTO) error
	Update(ctx context.Context, cfg configuration.UpdateComputerConfigurationDTO) error
	Delete(ctx context.Context, id string) error
}
