package stations

import (
	"container/list"
	"fmt"
	"os"
)

const (
	StationsField = iota
	ConnectionsField
)

func (graph *Graph) ParseTrains(line string, numTrains int) error {

	graph.NumTrains = numTrains
	graph.data.field = StationsField

	return nil
}

func (graph *Graph) ParseStations(line string, startStation string, endStation string, numTrains int) error {

	// Check if there are more than 10,000 Stations
	if len(graph.Stations) > 10000 {
		fmt.Fprintln(os.Stderr, "ERROR: maximum amount of stations is 10,000")
		os.Exit(0)
	}

	station, err := getStation(line, graph)
	if err != nil {
		return err
	}

	if graph.Stations[station] != nil {
		fmt.Fprintln(os.Stderr, "ERROR: found duplicate station:", station)
		os.Exit(0)
	}

	if station != "" {
		node := &Node{Parent: "T", Edges: make(map[string]byte)}
		graph.Start = startStation
		graph.End = endStation
		graph.Stations[station] = &node
		graph.Exits = list.New()
	} else if graph.Start != "" && graph.End != "" {
		graph.data.field = ConnectionsField
		graph.data.coordinates = nil
		return graph.ParseData(line, startStation, endStation, numTrains)
	}
	return nil
}

func (graph *Graph) ParseConnection(line string) error {
	station1, station2 := getConnection(line)
	if station1 == "" || station2 == "" {
		fmt.Fprintf(os.Stderr, "ERROR: invalid station\n")
		os.Exit(0)
		return fmt.Errorf("invalid station")
	}

	// Check if both stations exist in the graph
	node1, ok1 := graph.Stations[station1]
	if !ok1 {
		fmt.Fprintf(os.Stderr, "ERROR: Connection made with a station that does not exist: %s\n", station1)
		os.Exit(0)
	}

	node2, ok2 := graph.Stations[station2]
	if !ok2 {
		fmt.Fprintf(os.Stderr, "ERROR: Connection made with a station that does not exist: %s\n", station2)
		os.Exit(0)
	}

	// Check for existing connection
	if _, ok := (*node1).Edges[station2]; ok {
		fmt.Fprintf(os.Stderr, "ERROR: Duplicate connection found between %s and %s\n", station1, station2)
		os.Exit(0)
	}

	// Add connection if it doesn't exist
	(*node1).Edges[station2] = 1 // Assuming the byte value isn't significant, otherwise adjust as needed
	(*node2).Edges[station1] = 1

	// Update FromStart flag if one of the stations is the start station
	if station1 == graph.Start {
		(*node2).FromStart = true
	} else if station2 == graph.Start {
		(*node1).FromStart = true
	}

	return nil
}

func (graph *Graph) ParseData(line string, startStation string, endStation string, numTrains int) error {

	switch graph.data.field {
	case ConnectionsField:
		return graph.ParseConnection(line)
	case StationsField:

		graph.ParseStations(line, startStation, endStation, numTrains)

		return nil
	default:
		return fmt.Errorf("something went wrong while parsing")
	}
}
