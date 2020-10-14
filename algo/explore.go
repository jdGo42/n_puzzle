package main

import (
	"./goal_generator"
	"./manhattan"
	"fmt"
	//"math/rand"
)

var g_size = 3
var goalState = goalGenerator.Generator(g_size)

//var init_state = []int{2, 0, 3, 1}
var init_state = []int{0, 2, 7, 4, 1, 3, 8, 6, 5}

type position struct {
	state     []int
	cost      int
	heuristic int
	prev      int
}

func getHeuristic(state []int) int {
	return manhattanDistance.GetStateScore(g_size, state, goalState)
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
	for i := 0; i < len(closeList); i++ {
		if isSameState(pos.state, closeList[i].state) {
			return openList
		}
	}
	l := -1
	for i := 0; i < len(openList); i++ {
		if pos.heuristic >= openList[i].heuristic && l == -1 {
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
	//	fmt.Println(pos)
	//	fmt.Println(openList)
	if l == -1 {
		l = len(openList)
	}
	tmp := append(openList[:l], append([]position{pos}, openList[l:]...)...)
	//	fmt.Println(tmp, "\n")
	return tmp

}

func visitPosition(pos position, openList []position, closedList []position) []position {
	zeroIndex := getZeroIndex(pos.state)
	indexPrevPos := len(closedList) - 1

	state := make([]int, len(pos.state))
	var newPos position
	if zeroIndex/g_size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-g_size]
		state[zeroIndex-g_size] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex/g_size < g_size-1 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex+g_size]
		state[zeroIndex+g_size] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%g_size > 0 {
		copy(state, pos.state)
		state[zeroIndex] = state[zeroIndex-1]
		state[zeroIndex-1] = 0
		newPos = createPosition(state, pos.cost+1, indexPrevPos)
		openList = insertInOpenList(openList, closedList, newPos)
	}
	if zeroIndex%g_size < g_size-1 {
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
	tmp.heuristic = getHeuristic(state) + cost
	tmp.prev = prev
	return tmp
}

func rewind(pos position, closedList []position) {
	if pos.prev >= 0 {
		rewind(closedList[pos.prev], closedList)
	}
	fmt.Println(pos.state)
}

func Resolve(initial_state []int) {
	closedList := make([]position, 0)
	openList := make([]position, 0)
	var pos position

	start := createPosition(initial_state, 0, -1)
	openList = append(openList, start)
	i := 0
	for len(openList) != 0 {
		//	if i > 10 {
		//		return
		//	}
		pos = openList[len(openList)-1]
		openList = openList[:len(openList)-1]
		//fmt.Println(i, pos, len(openList))
		i++
		if isSameState(pos.state, goalState) {
			rewind(pos, closedList)
			return
		}
		closedList = append(closedList, pos)
		openList = visitPosition(pos, openList, closedList)
	}
	fmt.Println("Unsolvable puzzle")
}

func main() {
	Resolve(init_state)
}
