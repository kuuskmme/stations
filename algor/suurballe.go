package algor

import (
	"stations/stations"
)

func Suurballe(graph *stations.Graph) bool {
	if !Dijkstra(graph) {
		return false
	}
	CachePath(graph)
	return true
}
func CachePath(graph *stations.Graph) bool {
	var unsplit bool
	w := graph.End
	if graph.Stations[w] == nil { // Add nil check
		return false
	}

	v := (*graph.Stations[w]).EdgeIn
	// Check if graph.Exits already includes v
	for connection := graph.Exits.Front(); connection != nil; connection = connection.Next() {
		if connection.Value.(string) == v {
			return false
		}
	}
	graph.Exits.PushBack(v)
	for w != graph.Start {
		if graph.Stations[v] == nil || graph.Stations[w] == nil { // Add nil check
			return false
		}
		if (*graph.Stations[v]).Parent == w {
			if unsplit {
				unsplitNode(graph, w)
			}
			unsplit = true
			simultAssign(&w, &v, v, (*graph.Stations[v]).EdgeIn)
		} else {
			(*graph.Stations[w]).Parent = v
			splitNode(graph, w)
			unsplit = false
			simultAssign(&w, &v, v, (*graph.Stations[v]).EdgeOut)
		}
	}
	return true
}
func unsplitNode(graph *stations.Graph, v string) {
	(*graph.Stations[v]).Split = false
	(*graph.Stations[v]).Parent = "T"
}
func splitNode(graph *stations.Graph, v string) {
	if v != graph.Start && v != graph.End {
		(*graph.Stations[v]).Split = true
	}
}
func simultAssign(a, b *string, c, d string) {
	*a = c
	*b = d
}
