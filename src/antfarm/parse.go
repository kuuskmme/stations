package antfarm

import (
	"container/list"
	"fmt"
	"stations/config"
)

const (
	RoomsField = iota
	LinksField
)

func (graph *Graph) ParseAnts(line string, numTrains int) error {

	graph.Nants = numTrains
	graph.data.field = RoomsField

	return nil
}

func (graph *Graph) ParseRooms(line string, startStation string, endStation string, numTrains int) error {

	if isStart(line) && !graph.data.startFound {
		graph.data.startFound = true
	} else if isEnd(line) && !graph.data.endFound {
		graph.data.endFound = true
	} else {
		room, err := getRoom(line, graph)
		if err != nil {
			return err
		}
		if graph.Rooms[room] != nil {
			return fmt.Errorf("the %v room is duplicated", room)
		}
		if room != "" {
			node := &Node{Parent: "T", Edges: make(map[string]byte)}
			graph.Start = startStation
			graph.End = endStation
			graph.Rooms[room] = &node
			graph.Exits = list.New()
		} else if graph.Start != "" && graph.End != "" {
			graph.data.field = LinksField
			graph.data.coordinates = nil
			return graph.ParseData(line, startStation, endStation, numTrains)
		} else {
			return fmt.Errorf(config.ErrNoStart + " or " + config.ErrNoEnd)
		}
	}

	return nil
}

func (graph *Graph) ParseLinks(line string) error {
	room1, room2 := getLink(line)
	if room1 == "" && room2 == "" {
		return fmt.Errorf("invalid link")
	}
	if graph.Rooms[room1] == nil || graph.Rooms[room2] == nil {
		return fmt.Errorf("the link contains an unknown room: %v", line)
	}
	if room1 == room2 {
		return fmt.Errorf("the %v room is linked to itself: %v", room1, line)
	}
	node1 := *graph.Rooms[room1]
	node1.Edges[room2] = 1 // map's value won't be used
	node2 := *graph.Rooms[room2]
	node2.Edges[room1] = 1

	// Check if one of the rooms is the start room and update the connected room's FromStart property
	if room1 == graph.Start {
		node2.FromStart = true
	} else if room2 == graph.Start {
		node1.FromStart = true
	}

	return nil
}

func (graph *Graph) ParseData(line string, startStation string, endStation string, numTrains int) error {
	switch graph.data.field {
	case LinksField:
		return graph.ParseLinks(line)
	case RoomsField:
		return graph.ParseRooms(line, startStation, endStation, numTrains)
	default:
		return fmt.Errorf("something went wrong while parsing")
	}
}
