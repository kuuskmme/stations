package antfarm

import "container/list"

type Graph struct {
	Rooms      map[string]**Node
	Exits      *list.List
	Start, End string
	Nants      int
	data       parseInfo
}

type parseInfo struct {
	coordinates          map[[2]int]bool
	field                byte
	startFound, endFound bool
}

type Node struct {
	Edges             map[string]byte
	Parent            string
	EdgeIn, EdgeOut   string
	PriceIn, PriceOut int
	CostIn, CostOut   int
	Split             bool
	FromStart         bool // Added to track if coming directly from the start point
	DirectUsed        bool // Flag to indicate if the direct start-to-end path has been used this turn
}

func NewGraph() *Graph {
	return &Graph{Rooms: make(map[string]**Node)}
}

// Hypothetical function to update node state as an entity moves to it
func (n *Node) UpdateStateFromStart(previousNode *Node, graph *Graph) {
	// Dereference once to get *Node from **Node
	startNode := *graph.Rooms[graph.Start]

	// If the previous node is the start node, or if the previous node itself is coming from start
	if previousNode == startNode || previousNode.FromStart {
		n.FromStart = true
	} else {
		n.FromStart = false
	}
}
