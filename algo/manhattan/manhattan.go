package manhattanDistance

import (
	//"fmt"
	"strconv"
)

func abs(n int) int {
	y := n >> (strconv.IntSize - 1)
	return (n ^ y) - y
}

func GetStateScore(size int, state []int, goal []int) int {
	sum := 0
	l := size * size

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if state[i] == goal[j] {
				sum += abs(j/size-i/size) + abs(j%size-i%size)
				break
			}
		}
	}
	return sum
}

/*
func main() {
	size := 3
	state := []int{2, 7, 5, 8, 0, 1, 3, 6, 4}
	fmt.Println(ManhattanDistance(size, state))
}
*/
