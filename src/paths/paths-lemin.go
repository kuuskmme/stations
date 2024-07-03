package paths

import (
	"container/list"
	"fmt"
)

func Lemin(paths *Paths, nants int) {
	if paths.pathLen(0) == 1 {
		room := (*paths.AllPaths[0]).Front().Value
		for i := 1; i <= nants; i++ {
			fmt.Printf("T%d-%v ", i, room)
		}
		fmt.Println()
		return
	}
	paths.antsSplit(nants) //splits ants into each path
	var lastAnt int
	antNum, activeAnt := 1, 1
	m := make(map[int]*list.Element)
	for j := 0; j < paths.Nsteps; j++ {
		for k := activeAnt; k <= lastAnt; k++ {
			if val, ok := m[k]; ok && val != nil {
				fmt.Printf("T%d-%v ", k, val.Value)
				m[k] = val.Next()
			} else {
				activeAnt = k + 1
				delete(m, k)
			}
		}
		for i := 0; i < paths.Npaths; i++ {
			if antNum > nants {
				break
			}
			if paths.Assignment[i] <= 0 {
				continue
			} else {
				paths.Assignment[i]--
			}
			fmt.Printf("T%d-%v ", antNum, (*paths.AllPaths[i]).Front().Value)
			m[antNum] = (*paths.AllPaths[i]).Front().Next()
			antNum++
		}
		fmt.Println()
		lastAnt = antNum - 1
	}
}
