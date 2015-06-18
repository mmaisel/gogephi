package gephi

type Node struct {
	Entity
	Label string `json:"label"`
	Size  int    `json:"size"`
	Coordinates
	RGB
}

func NewNode(id string) (n *Node) {
	n = new(Node)
	n.ID = id
	n.Attributes = make(map[string]interface{})
	return
}

func (n *Node) GetAttributes() (a Attributes) {
	a = n.Attributes
	a["label"] = n.Label
	a["size"] = n.Size
	a["x"] = n.Coordinates.X
	a["y"] = n.Coordinates.Y
	a["z"] = n.Coordinates.Z
	a["r"] = n.RGB.Red
	a["b"] = n.RGB.Blue
	a["g"] = n.RGB.Green
	return
}
