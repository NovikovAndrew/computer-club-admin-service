package components

import "context"

type service struct {
	storage Storage
}

func NewComponentStorage(storage Storage) *service {
	return &service{
		storage: storage,
	}
}

func (s *service) GetOneCPU(ctx context.Context, id string) (*CPU, error) {
	panic("impl")
}
func (s *service) GetAllCPU(ctx context.Context, limit, offset int64) ([]CPU, error) {
	panic("impl")
}
func (s *service) CreateCPU(ctx context.Context, cpu CreateCPUDTO) error {
	panic("impl")
}
func (s *service) UpdateCPU(ctx context.Context, cpu UpdateCPUDTO) error {
	panic("impl")
}
func (s *service) DeleteCPU(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneGPU(ctx context.Context, id string) (*GPU, error) {
	panic("impl")
}
func (s *service) GetAllGPU(ctx context.Context, limit, offset int64) ([]GPU, error) {
	panic("impl")
}
func (s *service) CreateGPU(ctx context.Context, gpu CreateGPUDTO) error {
	panic("impl")
}
func (s *service) UpdateGPU(ctx context.Context, gpu UpdateGPUDTO) error {
	panic("impl")
}
func (s *service) DeleteGPU(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneRAM(ctx context.Context, id string) (*RAM, error) {
	panic("impl")
}
func (s *service) GetAllRAM(ctx context.Context, limit, offset int64) ([]RAM, error) {
	panic("impl")
}
func (s *service) CreateRAM(ctx context.Context, ram CreateRAMDTO) error {
	panic("impl")
}
func (s *service) UpdateRAM(ctx context.Context, ram UpdateRAMDTO) error {
	panic("impl")
}
func (s *service) DeleteRAM(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneMotherBoard(ctx context.Context, id string) (*MotherBoard, error) {
	panic("impl")
}
func (s *service) GetAllMotherBoard(ctx context.Context, limit, offset int64) ([]MotherBoard, error) {
	panic("impl")
}
func (s *service) CreateMotherBoard(ctx context.Context, motherBoard CreateMotherBoardDTO) error {
	panic("impl")
}
func (s *service) UpdateMotherBoard(ctx context.Context, motherBoard UpdateMotherBoardDTO) error {
	panic("impl")
}
func (s *service) DeleteMotherBoard(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOnePowerSupply(ctx context.Context, id string) (*PowerSupply, error) {
	panic("impl")
}
func (s *service) GetAllPowerSupply(ctx context.Context, limit, offset int64) ([]PowerSupply, error) {
	panic("impl")
}
func (s *service) CreatePowerSupply(ctx context.Context, powerSupply CreatePowerSupplyDTO) error {
	panic("impl")
}
func (s *service) UpdatePowerSupply(ctx context.Context, powerSupply UpdatePowerSupplyDTO) error {
	panic("impl")
}
func (s *service) DeletePowerSupply(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneMouse(ctx context.Context, id string) (*Mouse, error) {
	panic("impl")
}
func (s *service) GetAllMouse(ctx context.Context, limit, offset int64) ([]Mouse, error) {
	panic("impl")
}
func (s *service) CreateMouse(ctx context.Context, mouse CreateMouseDTO) error {
	panic("impl")
}
func (s *service) UpdateMouse(ctx context.Context, mouse UpdateMouseDTO) error {
	panic("impl")
}
func (s *service) DeleteMouse(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneKeyboard(ctx context.Context, id string) (*Keyboard, error) {
	panic("impl")
}
func (s *service) GetAllKeyboard(ctx context.Context, limit, offset int64) ([]Keyboard, error) {
	panic("impl")
}
func (s *service) CreateKeyboard(ctx context.Context, keyboard CreateKeyboardDTO) error {
	panic("impl")
}
func (s *service) UpdateKeyboard(ctx context.Context, keyboard UpdateKeyboardDTO) error {
	panic("impl")
}
func (s *service) DeleteKeyboard(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneHeadphones(ctx context.Context, id string) (*Headphones, error) {
	panic("impl")
}
func (s *service) GetAllHeadphones(ctx context.Context, limit, offset int64) ([]Headphones, error) {
	panic("impl")
}
func (s *service) CreateHeadphones(ctx context.Context, headphones CreateHeadphonesDTO) error {
	panic("impl")
}
func (s *service) UpdateHeadphones(ctx context.Context, headphones UpdateHeadphonesDTO) error {
	panic("impl")
}
func (s *service) DeleteHeadphones(ctx context.Context, id string) error {
	panic("impl")
}

func (s *service) GetOneMonitor(ctx context.Context, id string) (*Monitor, error) {
	panic("impl")
}
func (s *service) GetAllMonitor(ctx context.Context, limit, offset int64) ([]Monitor, error) {
	panic("impl")
}
func (s *service) CreateMonitor(ctx context.Context, monitor CreateMonitorDTO) error {
	panic("impl")
}
func (s *service) UpdateMonitor(ctx context.Context, monitor UpdateMonitorDTO) error {
	panic("impl")
}
func (s *service) DeleteMonitor(ctx context.Context, id string) error {
	panic("impl")
}
