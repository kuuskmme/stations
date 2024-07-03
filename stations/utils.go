package stations

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isCoord(x, y string, graph *Graph) bool {
	xn, err1 := strconv.Atoi(x)
	yn, err2 := strconv.Atoi(y)
	if xn < 0 {
		fmt.Fprintf(os.Stderr, "ERROR: found a negative X coordinate: %d\n", xn)
		os.Exit(0)
	}
	if yn < 0 {
		fmt.Fprintf(os.Stderr, "ERROR: found a negative Y coordinate: %d\n", yn)
		os.Exit(0)
	}
	if err1 != nil || err2 != nil {
		return false
	}
	// Check for existing coordinates
	if _, ok := graph.data.coordinates[[2]int{xn, yn}]; ok {
		fmt.Fprintln(os.Stderr, "ERROR: found stations with the same coordinates")
		os.Exit(0)
	}

	graph.data.coordinates[[2]int{xn, yn}] = true
	return true
}

func getStation(station string, graph *Graph) (string, error) {
	s := strings.Split(station, ",")
	if len(s) != 3 {
		return "", nil
	}

	if !isCoord(s[1], s[2], graph) {
		return "", fmt.Errorf("")
	}

	return s[0], nil
}

func getConnection(connection string) (string, string) {
	s := strings.Split(connection, "-")
	if len(s) != 2 {
		return "", ""
	}
	return s[0], s[1]
}
