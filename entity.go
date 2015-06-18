package gephi

type Attributes map[string]interface{}

type RGB struct {
	Red   float32 `json:"r"`
	Green float32 `json:"g"`
	Blue  float32 `json:"b"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Entity struct {
	ID string `json:"-"`
	Attributes
}
