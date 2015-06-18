package gephi

type Edge struct {
	Entity
	Label    string `json:"label"`
	Source   *Node
	Target   *Node
	Directed bool `json:"directed"`
	Weight   int  `json:"weight"`
	RGB
}

func NewEdge(id string, source *Node, target *Node) (e *Edge) {
	e = new(Edge)
	e.ID = id
	e.Attributes = make(map[string]interface{})
	e.Source = source
	e.Target = target
	return
}

func (e *Edge) GetAttributes() (a Attributes) {
	a = e.Attributes
	a["label"] = e.Label
	a["source"] = e.Source.ID
	a["target"] = e.Target.ID
	a["directed"] = e.Directed
	a["weight"] = e.Weight
	a["r"] = e.RGB.Red
	a["b"] = e.RGB.Blue
	a["g"] = e.RGB.Green
	return
}
