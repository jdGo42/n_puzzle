package main

import (
	"fmt"
)

var size = 2

type position struct {
	state     []int
	cost      int
	heuristic int
	prev      int
}

var g_i = 0

func getHeuristic(state []int) int {
	g_i++
	if g_i >= 5 {
		return 0
	} else {
		return 1
	}
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

func stateCmp(s1 []int, s2 []int) int {
	i := 0
	for i < len(s1) {
		if s1[i] != s2[i] {
			return 1
		}
		i++
	}
	return 0
}

func insertInOpenList(openList []position, closeList []position, pos position) []position {
	i := 0
	l := 0
	for i < len(closeList) {
		if stateCmp(pos.state, closeList[i].state) == 0 {
			return openList
		}
		i++
	}
	i = 0
	for i < len(openList) {
		if stateCmp(pos.state, openList[i].state) == 0 && openList[i].cost <= pos.cost {
			return openList
		}
		if pos.heuristic < openList[i].heuristic && l == 0 {
			l = i
		}
		i++
	}
	return append(openList[:l], append([]position{pos}, openList[l:]...)...)
}

func visitPosition(pos position, openList []position, closedList []position) []position {
	zeroIndex := getZeroIndex(pos.state)
	indexPrevPos := len(closedList) - 1

	state := make([]int, len(pos.state))
	var newPos position
	if zeroIndex/size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-size]
		state[zeroIndex-size] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex/size < size-1 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex+size]
		state[zeroIndex+size] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-1]
		state[zeroIndex-1] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%size < size-1 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex+1]
		state[zeroIndex+1] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	return openList
}

func createPosition(state []int, cost int, prev int) position {
	var tmp position
	tmp.state = make([]int, len(state))
	copy(tmp.state, state)
	tmp.cost = cost
	tmp.heuristic = cost + getHeuristic(state)
	tmp.prev = prev
	return tmp
}

func Resolve(initial_state []int) {
	var closedList []position
	var openList []position
	var pos position

	start := createPosition(initial_state, 0, -1)
	openList = append(openList, start)
	for len(openList) != 0 {
		pos = openList[len(openList)-1]
		if stateCmp(pos.state, []int{0, 1, 2, 3}) == 0 {
			break
		}
		openList = openList[:len(openList)-1]
		closedList = append(closedList, pos)
		openList = visitPosition(pos, openList, closedList)
		fmt.Println(closedList)
		fmt.Println(openList)
	}
	fmt.Println(pos)
}

func main() {
	init_state := []int{1, 3, 0, 2}
	Resolve(init_state)
}
