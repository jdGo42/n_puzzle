package resolve

import (
	"./goal_generator"
	"./manhattan"
	"fmt"
)

type position struct {
	state     []int
	cost      int
	heuristic int
	prev      int
}

func getHeuristic(size int, state []int, goalState []int) int {
	return manhattanDistance.GetStateScore(size, state, goalState)
}

func getZeroIndex(state []int) int {
	zeroIndex := 0
	for zeroIndex < len(state) {
		if state[zeroIndex] == 0 {
			break
		}
		zeroIndex++
	}
	return zeroIndex
}

func isSameState(s1 []int, s2 []int) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func insertInOpenList(openList []position, closeList []position, pos position) []position {
	for i := len(closeList) - 1; i >= 0; i-- {
		if isSameState(pos.state, closeList[i].state) {
			return openList
		}
	}
	l := -1
	for i := len(openList) - 1; i >= 0; i-- {
		if pos.heuristic < openList[i].heuristic && l == -1 {
			l = i
		}
		if isSameState(pos.state, openList[i].state) {
			if openList[i].cost <= pos.cost {
				return openList
			}
			copy(openList[i:], openList[i+1:])
			openList = openList[:len(openList)-1]
		}
	}
	if l != len(openList) {
		l++
		return append(openList[:l], append([]position{pos}, openList[l:]...)...)
	} else {
		return append(openList, pos)
	}

}

func visitPosition(size int, goalState []int, pos position, openList []position, closedList []position) []position {
	zeroIndex := getZeroIndex(pos.state)
	indexPrevPos := len(closedList) - 1

	state := make([]int, len(pos.state))
	var newPos position
	if zeroIndex/size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-size]
		state[zeroIndex-size] = 0
		newPos = createPosition(size, state, goalState, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex/size < size-1 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex+size]
		state[zeroIndex+size] = 0
		newPos = createPosition(size, state, goalState, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-1]
		state[zeroIndex-1] = 0
		newPos = createPosition(size, state, goalState, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%size < size-1 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex+1]
		state[zeroIndex+1] = 0
		newPos = createPosition(size, state, goalState, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	return openList
}

func createPosition(size int, state []int, goalState []int, cost int, prev int) position {
	var tmp position
	tmp.state = make([]int, len(state))
	copy(tmp.state, state)
	tmp.cost = cost
	tmp.heuristic = getHeuristic(size, state, goalState) + cost
	tmp.prev = prev
	return tmp
}

func rewind(pos position, closedList []position) {
	if pos.prev >= 0 {
		rewind(closedList[pos.prev], closedList)
	}
	fmt.Println(pos.state)
}

func Resolve(size int, initial_state []int) {
	closedList := make([]position, 0, 1024)
	openList := make([]position, 0, 1024)
	goalState := goalGenerator.Generator(size)
	var pos position

	start := createPosition(size, initial_state, goalState, 0, -1)
	openList = append(openList, start)
	i := 0
	for len(openList) != 0 {
		pos = openList[len(openList)-1]
		openList = openList[:len(openList)-1]
		i++
		if isSameState(pos.state, goalState) {
			rewind(pos, closedList)
			return
		}
		closedList = append(closedList, pos)
		openList = visitPosition(size, goalState, pos, openList, closedList)
	}
	fmt.Println("Unsolvable puzzle")
}

/*
func main() {
	Resolve(init_state)
}
*/
