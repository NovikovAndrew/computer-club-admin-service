package components

// CPU

type CreateCPUDTO struct {
	Name      string `json:"name"`
	Frequency int64  `json:"frequency"`
}

type UpdateCPUDTO struct {
	Name      string `json:"name"`
	Frequency int64  `json:"frequency"`
}

// GPU

type CreateGPUDTO struct {
	Name string `json:"name"`
}

type UpdateGPUDTO struct {
	Name string `json:"name"`
}

// RAM

type CreateRAMDTO struct {
	Name string `json:"name"`
	Size int32  `json:"size"`
}

type UpdateRAMDTO struct {
	Name string `json:"name"`
	Size int32  `json:"size"`
}

// MotherBoard

type CreateMotherBoardDTO struct {
	Name string `json:"name"`
}

type UpdateMotherBoardDTO struct {
	Name string `json:"name"`
}

// PowerSupply

type CreatePowerSupplyDTO struct {
	Name  string `json:"name"`
	Power int32  `json:"power"`
}

type UpdatePowerSupplyDTO struct {
	Name  string `json:"name"`
	Power int32  `json:"power"`
}

// Mouse

type CreateMouseDTO struct {
	Name string `json:"name"`
}

type UpdateMouseDTO struct {
	Name string `json:"name"`
}

// Keyboard

type CreateKeyboardDTO struct {
	Name string `json:"name"`
}

type UpdateKeyboardDTO struct {
	Name string `json:"name"`
}

// Headphones

type CreateHeadphonesDTO struct {
	Name string `json:"name"`
}

type UpdateHeadphonesDTO struct {
	Name string `json:"name"`
}

// Monitor

type CreateMonitorDTO struct {
	Name string `json:"name"`
}

type UpdateMonitorDTO struct {
	Name string `json:"name"`
}
