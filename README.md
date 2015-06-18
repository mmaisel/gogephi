# Golang Gephi Streaming API Client 
This is a golang client for the Gephi Streaming API.

# Usage

```
	g := gephi.NewClient("localhost:8080", "workspace0")

	a := gephi.NewNode("A")
	b := gephi.NewNode("B")
	b.Attributes["Foo"] = "Bar"
	b.Size = 100
	e := gephi.NewEdge("A --- B", &a, &b)
	e1 := gephi.NewEdge("A --> B", &a, &b)
	e1.Directed = true

	g.AddNode(a)
	g.AddNode(b)
	g.AddEdge(e)
	g.AddEdge(e1)

	g.Commit()
```
