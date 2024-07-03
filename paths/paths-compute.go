package paths

import (
	"container/list"
	"sort"
	"stations/algor"
	"stations/stations"
)

func PathsCompute(graph *stations.Graph) *Paths {
	var pathsOld *Paths
	var pathsNew *Paths
	if pathsOld = PathsGetNext(graph); pathsOld == nil {
		return nil
	}
	nPaths := 1
	for nPaths < graph.NumTrains {
		if pathsNew = PathsGetNext(graph); pathsNew == nil {
			break
		}
		if pathsNew.Nsteps < pathsOld.Nsteps {
			pathsOld = pathsNew
		}
		nPaths++
	}
	return pathsOld
}

func PathsGetNext(graph *stations.Graph) *Paths {
	if !algor.Suurballe(graph) {
		return nil
	}
	return PathsFromGraph(graph)
}

func PathsFromGraph(graph *stations.Graph) *Paths {
	paths := new(Paths)
	paths.Npaths = graph.Exits.Len()
	paths.AllPaths = make([]**list.List, paths.Npaths)
	i := 0
	for connection := graph.Exits.Front(); connection != nil; connection = connection.Next() {
		p := unrollPath(graph, connection.Value.(string))
		paths.AllPaths[i] = &p
		i++
	}
	sort.Slice(paths.AllPaths, func(i, j int) bool { return (*paths.AllPaths[i]).Len() < (*paths.AllPaths[j]).Len() })
	paths.Nsteps = paths.calcSteps(graph.NumTrains)
	return paths
}

func unrollPath(graph *stations.Graph, v string) *list.List {
	path := list.New()
	path.PushFront(graph.End)
	for v != graph.Start {
		path.PushFront(v)
		v = (*graph.Stations[v]).Parent
	}
	return path
}
