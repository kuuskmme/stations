package paths

import (
	"container/list"
	"fmt"
)

func PrintPaths(paths *Paths, nTrains int) {
	paths.TrainsSplit(nTrains) //splits trains into each path
	var lastTrain int
	trainNum, activeTrain := 1, 1
	m := make(map[int]*list.Element)
	for j := 0; j < paths.Nsteps; j++ {
		for k := activeTrain; k <= lastTrain; k++ {
			if val, ok := m[k]; ok && val != nil {
				fmt.Printf("T%d-%v ", k, val.Value)
				m[k] = val.Next()
			} else {
				activeTrain = k + 1
				delete(m, k)
			}
		}
		for i := 0; i < paths.Npaths; i++ {
			if trainNum > nTrains {
				break
			}
			if paths.Assignment[i] <= 0 {
				continue
			} else {
				paths.Assignment[i]--
			}
			fmt.Printf("T%d-%v ", trainNum, (*paths.AllPaths[i]).Front().Value)
			m[trainNum] = (*paths.AllPaths[i]).Front().Next()
			trainNum++
		}
		fmt.Println()
		lastTrain = trainNum - 1
	}
}
