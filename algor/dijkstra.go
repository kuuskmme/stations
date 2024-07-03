package algor

import (
	"container/heap"
	"stations/stations"
)

const MaxNNodes = 100000

func Dijkstra(graph *stations.Graph) bool {
	pq := make(PriorityQueue, 0, 100)
	var v string
	GraphReset(graph)

	if graph.Stations[graph.Start] != nil {
		heap.Push(&pq, &Node{v: 0, station: graph.Start})
	}

	for pq.Len() > 0 {
		poppedNode := heap.Pop(&pq).(*Node)
		if poppedNode == nil || graph.Stations[poppedNode.station] == nil {
			continue
		}
		v = poppedNode.station

		for w := range (*graph.Stations[v]).Edges {
			if graph.Stations[w] != nil { // Ensure the connected node exists
				RelaxEdge(graph, &pq, v, w)
			}
		}
	}

	SetPrices(graph) // Ensure this function only operates on valid nodes
	if endNode := graph.Stations[graph.End]; endNode != nil {
		return (*endNode).EdgeIn != "T"
	}
	return false
}

func GraphReset(graph *stations.Graph) {
	var node *stations.Node
	for _, value := range graph.Stations {
		node = *value
		node.EdgeIn = "T"
		node.EdgeOut = "T"
		node.CostIn = MaxNNodes
		node.CostOut = MaxNNodes
	}
	node = *graph.Stations[graph.Start]
	node.CostIn = 0
	node.CostOut = 0
}

func RelaxEdge(graph *stations.Graph, pq *PriorityQueue, v, w string) {
	if graph.Stations[v] == nil || graph.Stations[w] == nil {
		return
	}
	nodeV := *graph.Stations[v]
	nodeW := *graph.Stations[w]
	if v == graph.End || w == graph.Start || nodeW.Parent == v {
		return
	}
	if nodeV.Parent == w && nodeV.CostIn < MaxNNodes && (1+nodeW.CostOut > nodeV.CostIn+nodeV.PriceIn-nodeW.PriceOut) {
		nodeW.EdgeOut = v
		nodeW.CostOut = nodeV.CostIn - 1 + nodeV.PriceIn - nodeW.PriceOut
		heap.Push(pq, &Node{v: nodeW.CostOut, station: w})
		RelaxHiddenEdge(graph, pq, w)
	} else if nodeV.Parent != w && nodeV.CostOut < MaxNNodes && -1+nodeW.CostIn > nodeV.CostOut+nodeV.PriceOut-nodeW.PriceIn {
		nodeW.EdgeIn = v
		nodeW.CostIn = nodeV.CostOut + 1 + nodeV.PriceOut - nodeW.PriceIn
		heap.Push(pq, &Node{v: nodeW.CostIn, station: w})
		RelaxHiddenEdge(graph, pq, w)
	}
}

func RelaxHiddenEdge(graph *stations.Graph, pq *PriorityQueue, w string) {
	node := *graph.Stations[w]
	if node.Split && node.CostIn > node.CostOut+node.PriceOut-node.PriceIn && w != graph.Start {
		node.EdgeIn = node.EdgeOut
		node.CostIn = node.CostOut + node.PriceOut - node.PriceIn
		if node.CostIn != node.CostOut {
			heap.Push(pq, &Node{v: node.CostIn, station: w})
		}
	}
	if !node.Split && node.CostOut > node.CostIn+node.PriceIn-node.PriceOut && w != graph.End {
		node.EdgeOut = node.EdgeIn
		node.CostOut = node.CostIn + node.PriceIn - node.PriceOut
		if node.CostIn != node.CostOut {
			heap.Push(pq, &Node{v: node.CostOut, station: w})
		}
	}
}

func SetPrices(graph *stations.Graph) {
	for _, value := range graph.Stations {
		if value != nil { // Add nil check
			node := *value
			node.PriceIn = node.CostIn
			node.PriceOut = node.CostOut
		}
	}
}
