package configuration

import (
	"context"
)

type Storage interface {
	GetOne(ctx context.Context, id string) (*ComputerConfiguration, error)
	GetAllByComputerID(ctx context.Context, id string) ([]ComputerConfiguration, error)
	Create(ctx context.Context, cfg ComputerConfiguration) error
	Update(ctx context.Context, cfg ComputerConfiguration) error
	Delete(ctx context.Context, id string) error
}
