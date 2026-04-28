package graph

type Edge struct {
	to     int
	weight float64
}

type Graph struct {
	adjList map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(from, to int, weight float64) {
	g.adjList[from] = append(g.adjList[from], Edge{to, weight})
}


