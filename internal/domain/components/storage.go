package components

import "context"

type Storage interface {
	GetOneCPU(ctx context.Context, id string) (*CPU, error)
	GetAllCPU(ctx context.Context, limit, offset int64) ([]CPU, error)
	CreateCPU(ctx context.Context, cpu CPU) error
	UpdateCPU(ctx context.Context, cpu CPU) error
	DeleteCPU(ctx context.Context, id string) error

	GetOneGPU(ctx context.Context, id string) (*GPU, error)
	GetAllGPU(ctx context.Context, limit, offset int64) ([]GPU, error)
	CreateGPU(ctx context.Context, gpu GPU) error
	UpdateGPU(ctx context.Context, gpu GPU) error
	DeleteGPU(ctx context.Context, id string) error

	GetOneRAM(ctx context.Context, id string) (*RAM, error)
	GetAllRAM(ctx context.Context, limit, offset int64) ([]RAM, error)
	CreateRAM(ctx context.Context, ram RAM) error
	UpdateRAM(ctx context.Context, ram RAM) error
	DeleteRAM(ctx context.Context, id string) error

	GetOneMotherBoard(ctx context.Context, id string) (*MotherBoard, error)
	GetAllMotherBoard(ctx context.Context, limit, offset int64) ([]MotherBoard, error)
	CreateMotherBoard(ctx context.Context, motherBoard MotherBoard) error
	UpdateMotherBoard(ctx context.Context, motherBoard MotherBoard) error
	DeleteMotherBoard(ctx context.Context, id string) error

	GetOnePowerSupply(ctx context.Context, id string) (*PowerSupply, error)
	GetAllPowerSupply(ctx context.Context, limit, offset int64) ([]PowerSupply, error)
	CreatePowerSupply(ctx context.Context, powerSupply PowerSupply) error
	UpdatePowerSupply(ctx context.Context, powerSupply PowerSupply) error
	DeletePowerSupply(ctx context.Context, id string) error

	GetOneMouse(ctx context.Context, id string) (*Mouse, error)
	GetAllMouse(ctx context.Context, limit, offset int64) ([]Mouse, error)
	CreateMouse(ctx context.Context, mouse Mouse) error
	UpdateMouse(ctx context.Context, mouse Mouse) error
	DeleteMouse(ctx context.Context, id string) error

	GetOneKeyboard(ctx context.Context, id string) (*Keyboard, error)
	GetAllKeyboard(ctx context.Context, limit, offset int64) ([]Keyboard, error)
	CreateKeyboard(ctx context.Context, keyboard Keyboard) error
	UpdateKeyboard(ctx context.Context, keyboard Keyboard) error
	DeleteKeyboard(ctx context.Context, id string) error

	GetOneHeadphones(ctx context.Context, id string) (*Headphones, error)
	GetAllHeadphones(ctx context.Context, limit, offset int64) ([]Headphones, error)
	CreateHeadphones(ctx context.Context, headphones Headphones) error
	UpdateHeadphones(ctx context.Context, headphones Headphones) error
	DeleteHeadphones(ctx context.Context, id string) error

	GetOneMonitor(ctx context.Context, id string) (*Monitor, error)
	GetAllMonitor(ctx context.Context, limit, offset int64) ([]Monitor, error)
	CreateMonitor(ctx context.Context, monitor Monitor) error
	UpdateMonitor(ctx context.Context, monitor Monitor) error
	DeleteMonitor(ctx context.Context, id string) error
}
