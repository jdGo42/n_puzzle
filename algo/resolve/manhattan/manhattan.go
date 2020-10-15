package manhattanDistance

const intSize = 32 << (^uint(0) >> 63)

func abs(n int) int {
	y := n >> (intSize - 1)
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
