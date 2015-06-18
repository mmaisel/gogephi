package gephi

import (
	"testing"
)

var (
	g         = NewGephi("", "")
	testNodes = []Node{
		Node{Entity: Entity{ID: "A", Attributes: make(map[string]interface{})}, Label: "Test A", Size: 10},
		Node{Entity: Entity{ID: "B", Attributes: make(map[string]interface{})}, Label: "Test B", Size: 10},
		Node{Entity: Entity{ID: "C", Attributes: make(map[string]interface{})}, Label: "Test C", Size: 10},
	}
	testEdges = []Edge{
		Edge{Entity: Entity{ID: "A--B", Attributes: make(map[string]interface{})}, Label: "A--B", Source: &testNodes[0], Target: &testNodes[1], Directed: false, Weight: 10},
		Edge{Entity: Entity{ID: "A->C", Attributes: make(map[string]interface{})}, Label: "A->C", Source: &testNodes[0], Target: &testNodes[2], Directed: true, Weight: 5},
	}
)

func TestAdd(t *testing.T) {
	for _, n := range testNodes {
		g.AddNode(n)
	}
	for _, e := range testEdges {
		g.AddEdge(e)
	}
	g.Commit()
}

func TestChange(t *testing.T) {
	for _, n := range testNodes {
		n.Attributes["foo"] = "bar"
		g.ChangeNode(n)
	}
	for _, e := range testEdges {
		e.Attributes["bar"] = "blah"
		g.ChangeEdge(e)
	}
	g.Commit()
}

func TestDelete(t *testing.T) {
	for _, n := range testNodes {
		g.DeleteNode(n)
	}
	for _, e := range testEdges {
		g.DeleteEdge(e)
	}
	g.Commit()
}

func TestNew(t *testing.T) {
	k := *NewNode("K")
	j := *NewNode("J")
	e := *NewEdge("e1", &k, &j)
	g.AddNode(j)
	g.AddNode(k)
	g.AddEdge(e)
	g.Commit()
}
