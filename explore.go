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

// if !(newPos is in closedList || (newPos is in openList with cost < newCost))
// append to openList

func stateCmp(s1 []int, s2 []int) int {
	i := 0
	for i < len(s1) {
		if s1[i] != s2[2] {
			return 1
		}
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
	}
	for i < len(openList) {
		if stateCmp(pos.state, openList[i].state) == 0 && openList[i].cost <= pos.cost {
			return openList
		}
		if pos.heuristic < openList[i].heuristic && l == 0 {
			l = i
		}
	}
	return append(openList[:l], append([]position{pos}, openList[l:]...)...)
}

func visitPosition(pos position, openList []position, closedList []position) {
	closedList = append(closedList, pos)
	zeroIndex := getZeroIndex(pos.state)
	posPrev := len(closedList) - 1

	state := make([]int, len(pos.state))
	var newPos position
	if zeroIndex/size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-size]
		state[zeroIndex-size] = 0
		newPos = createPosition(state, pos.cost+1, posPrev)
		openList = insertInOpenList(openList, closeList, newPos)
	}
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

	start := createPosition(initial_state, -1, -1)
	openList = append(openList, start)

	fmt.Println(closedList)
	fmt.Println(openList)
}

func main() {
	init_state := []int{0, 1, 2, 3}
	Resolve(init_state)
}
