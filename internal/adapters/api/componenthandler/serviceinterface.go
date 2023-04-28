package componenthandler

import (
	"context"

	"github.com/NovikovAndrew/computer-club-admin-service.git/internal/domain/components"
)

type ComponentService interface {
	GetOneCPU(ctx context.Context, id string) (*components.CPU, error)
	GetAllCPU(ctx context.Context, limit, offset int64) ([]components.CPU, error)
	CreateCPU(ctx context.Context, cpu components.CreateCPUDTO) error
	UpdateCPU(ctx context.Context, cpu components.UpdateCPUDTO) error
	DeleteCPU(ctx context.Context, id string) error

	GetOneGPU(ctx context.Context, id string) (*components.GPU, error)
	GetAllGPU(ctx context.Context, limit, offset int64) ([]components.GPU, error)
	CreateGPU(ctx context.Context, gpu components.CreateGPUDTO) error
	UpdateGPU(ctx context.Context, gpu components.UpdateGPUDTO) error
	DeleteGPU(ctx context.Context, id string) error

	GetOneRAM(ctx context.Context, id string) (*components.RAM, error)
	GetAllRAM(ctx context.Context, limit, offset int64) ([]components.RAM, error)
	CreateRAM(ctx context.Context, ram components.CreateRAMDTO) error
	UpdateRAM(ctx context.Context, ram components.UpdateRAMDTO) error
	DeleteRAM(ctx context.Context, id string) error

	GetOneMotherBoard(ctx context.Context, id string) (*components.MotherBoard, error)
	GetAllMotherBoard(ctx context.Context, limit, offset int64) ([]components.MotherBoard, error)
	CreateMotherBoard(ctx context.Context, motherBoard components.CreateMotherBoardDTO) error
	UpdateMotherBoard(ctx context.Context, motherBoard components.UpdateMotherBoardDTO) error
	DeleteMotherBoard(ctx context.Context, id string) error

	GetOnePowerSupply(ctx context.Context, id string) (*components.PowerSupply, error)
	GetAllPowerSupply(ctx context.Context, limit, offset int64) ([]components.PowerSupply, error)
	CreatePowerSupply(ctx context.Context, powerSupply components.CreatePowerSupplyDTO) error
	UpdatePowerSupply(ctx context.Context, powerSupply components.UpdatePowerSupplyDTO) error
	DeletePowerSupply(ctx context.Context, id string) error

	GetOneMouse(ctx context.Context, id string) (*components.Mouse, error)
	GetAllMouse(ctx context.Context, limit, offset int64) ([]components.Mouse, error)
	CreateMouse(ctx context.Context, mouse components.CreateMouseDTO) error
	UpdateMouse(ctx context.Context, mouse components.UpdateMouseDTO) error
	DeleteMouse(ctx context.Context, id string) error

	GetOneKeyboard(ctx context.Context, id string) (*components.Keyboard, error)
	GetAllKeyboard(ctx context.Context, limit, offset int64) ([]components.Keyboard, error)
	CreateKeyboard(ctx context.Context, keyboard components.CreateKeyboardDTO) error
	UpdateKeyboard(ctx context.Context, keyboard components.UpdateKeyboardDTO) error
	DeleteKeyboard(ctx context.Context, id string) error

	GetOneHeadphones(ctx context.Context, id string) (*components.Headphones, error)
	GetAllHeadphones(ctx context.Context, limit, offset int64) ([]components.Headphones, error)
	CreateHeadphones(ctx context.Context, headphones components.CreateHeadphonesDTO) error
	UpdateHeadphones(ctx context.Context, headphones components.UpdateHeadphonesDTO) error
	DeleteHeadphones(ctx context.Context, id string) error

	GetOneMonitor(ctx context.Context, id string) (*components.Monitor, error)
	GetAllMonitor(ctx context.Context, limit, offset int64) ([]components.Monitor, error)
	CreateMonitor(ctx context.Context, monitor components.CreateMonitorDTO) error
	UpdateMonitor(ctx context.Context, monitor components.UpdateMonitorDTO) error
	DeleteMonitor(ctx context.Context, id string) error
}
