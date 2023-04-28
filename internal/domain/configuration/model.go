package configuration

import "github.com/NovikovAndrew/computer-club-admin-service.git/internal/domain/components"

type ComputerConfiguration struct {
	ID             string                 `json:"id" bson:"_id,omitempty"`
	ComputerClubID string                 `json:"computer_club_id" bson:"computer_club_id"`
	CPU            components.CPU         `json:"cpu" bson:"cpu"`
	GPU            components.GPU         `json:"gpu" bson:"gpu"`
	RAM            components.RAM         `json:"ram" bson:"ram"`
	MotherBoard    components.MotherBoard `json:"motherboard" bson:"motherboard"`
	PowerSupply    components.PowerSupply `json:"powersupply" bson:"powersupply"`
	Mouse          components.Mouse       `json:"mouse" bson:"mouse"`
	Keyboard       components.Keyboard    `json:"keyboard" bson:"keyboard"`
	Headphones     components.Headphones  `json:"headphones" bson:"headphones"`
}
