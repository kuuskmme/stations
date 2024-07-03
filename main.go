package main

import (
	"bufio"
	"fmt"
	"os"
	"stations/config"
	"stations/src/antfarm"
	"stations/src/paths"
	"strconv"
	"strings"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadData(filename string, startStation string, endStation string, numTrains int) (*antfarm.Graph, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf(config.ErrFileIssue)
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf(config.ErrFileIssue)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := antfarm.NewGraph()

	stationsSectionFound := false
	connectionsSectionFound := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		line = strings.ReplaceAll(line, " ", "")

		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
			line = strings.TrimSpace(line)
		}

		if line == "" {
			continue
		}

		if line == "stations:" {
			stationsSectionFound = true
			continue
		}

		if line == "connections:" {
			connectionsSectionFound = true
			continue
		}

		graph.ParseData(line, startStation, endStation, numTrains)
		graph.ParseAnts(line, numTrains)
	}

	if !stationsSectionFound {
		return nil, fmt.Errorf("NO STATIONS SECTION")
	}
	if !connectionsSectionFound {
		return nil, fmt.Errorf("NO CONNECTION SECTION")
	}
	return graph, nil
}

func main() {
	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrains, _ := strconv.Atoi(os.Args[4])

	graph, err := ReadData(filePath, startStation, endStation, numTrains)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	allPaths := paths.PathsCompute(graph)
	if allPaths == nil {
		fmt.Print(config.ErrNoPaths)
		os.Exit(1)
	}
	paths.Lemin(allPaths, graph.Nants)
}
