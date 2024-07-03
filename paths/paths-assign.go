package paths

import "container/list"

type Paths struct {
	Npaths, Nsteps int
	AllPaths       []**list.List
	Assignment     []int //groups of trains for each path
}

func (paths *Paths) pathLen(i int) int {
	return (*paths.AllPaths[i]).Len()
}

func (paths *Paths) calcSteps(nTrains int) int {
	l := len(paths.AllPaths) - 1
	shortest := paths.pathLen(0)
	longest := paths.pathLen(l)
	var sum int
	for i := 0; i < paths.Npaths; i++ {
		sum += longest - paths.pathLen(i)
	}
	numOfTrains := longest - shortest + (nTrains-sum)/paths.Npaths
	rem := (nTrains - sum) % paths.Npaths
	if rem > 0 {
		numOfTrains++
	}
	return shortest + numOfTrains - 1
}

func (paths *Paths) TrainsSplit(nTrains int) {
	paths.Assignment = make([]int, paths.Npaths)
	l := len(paths.AllPaths) - 1
	longest := paths.pathLen(l)
	var sum int
	for i := 0; i < paths.Npaths; i++ {
		sum += longest - paths.pathLen(i)
	}
	fn := float32(nTrains-sum) / float32(paths.Npaths)
	remSteps := (fn - float32(int(fn))) * float32(paths.Npaths)
	for i := 0; i < paths.Npaths; i++ {
		paths.Assignment[i] = longest - paths.pathLen(i) + int(fn)
		if remSteps > 0 {
			paths.Assignment[i]++
			remSteps--
		}
	}
}
