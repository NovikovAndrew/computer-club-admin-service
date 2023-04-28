package components

// MARK: - Computer Components

type CPU struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name"`
	Frequency int64  `json:"frequency" bson:"frequency"`
}

type GPU struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type RAM struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Size int32  `json:"size" bson:"size"`
}

type MotherBoard struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type PowerSupply struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Power int32  `json:"power" bson:"power"`
}

// MARK: - Devices

type Mouse struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type Keyboard struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type Headphones struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

type Monitor struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
