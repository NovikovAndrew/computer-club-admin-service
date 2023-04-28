package configuration

type CreateComputerConfigurationDTO struct {
	ComputerClubID string `json:"computer_club_id"`
	CPUID          string `json:"cpu_id"`
	GPUID          string `json:"gpu_id"`
	RAMID          string `json:"ram_id"`
	MotherBoardID  string `json:"motherboard_id"`
	PowerSupplyID  string `json:"powersupply_id"`
	MouseID        string `json:"mouse_id"`
	KeyboardID     string `json:"keyboard_id"`
	HeadphonesID   string `json:"headphones_id"`
}

type UpdateComputerConfigurationDTO struct {
	ComputerClubID string `json:"computer_club_id"`
	CPUID          string `json:"cpu_id"`
	GPUID          string `json:"gpu_id"`
	RAMID          string `json:"ram_id"`
	MotherBoardID  string `json:"motherboard_id"`
	PowerSupplyID  string `json:"powersupply_id"`
	MouseID        string `json:"mouse_id"`
	KeyboardID     string `json:"keyboard_id"`
	HeadphonesID   string `json:"headphones_id"`
}
