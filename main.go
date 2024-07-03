package main

import (
	"bufio"
	"fmt"
	"os"
	"stations/paths"
	"stations/stations"
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

func ReadData(filename string, startStation string, endStation string, numTrains int) *stations.Graph {
	if !FileExists(filename) {
		fmt.Fprintln(os.Stderr, "ERROR: file does not exist")
		return nil // Exit the function, indicating an error occurred by returning nil
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: could not open the file")
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := stations.NewGraph()

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

		graph.ParseTrains(line, numTrains)
	}

	if !stationsSectionFound {
		fmt.Fprintf(os.Stderr, "ERROR: no stations section\n")
		return nil
	}
	if !connectionsSectionFound {
		fmt.Fprintf(os.Stderr, "ERROR: no connections section\n")
		return nil
	}
	// Validate stations after fully reading the file
	if _, exists := graph.Stations[startStation]; !exists {
		fmt.Fprintf(os.Stderr, "ERROR: start station does not exist: %s\n", startStation)
		return nil
	}
	if _, exists := graph.Stations[endStation]; !exists {
		fmt.Fprintf(os.Stderr, "ERROR: end station does not exist: %s\n", endStation)
		return nil
	}
	return graph
}

func main() {
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "ERROR: too few command line arguments \nUsage: go run . <file path> <start station> <end station> <number of trains>\n")
		return
	}
	if len(os.Args) > 5 {
		fmt.Fprintf(os.Stderr, "ERROR: too many command line arguments \nUsage: go run . <file path> <start station> <end station> <number of trains>\n")
		return
	}
	filePath := os.Args[1]
	startStation := os.Args[2]
	endStation := os.Args[3]
	numTrains, err := strconv.Atoi(os.Args[4])
	if err != nil || numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "ERROR: Number of trains must be a positive integer and %d is not a positive integer\n", numTrains)
		return
	}
	if startStation == endStation {
		fmt.Fprintf(os.Stderr, "ERROR: start and end stations are the exact same: %s-%s\n", startStation, endStation)
		return
	}

	graph := ReadData(filePath, startStation, endStation, numTrains)
	if graph == nil {
		return
	}

	allPaths := paths.PathsCompute(graph)
	if allPaths == nil {
		fmt.Fprintln(os.Stderr, "ERROR: no path exist between the start and end stations")
		return
	}

	paths.PrintPaths(allPaths, graph.NumTrains)
}
