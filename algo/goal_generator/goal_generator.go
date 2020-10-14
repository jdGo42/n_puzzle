package goalGenerator

func Generator(size int) []int {
	goal := make([]int, size*size, size*size)
	cur := 1
	x := 0
	ix := 1
	y := 0
	iy := 0
	for true {
		goal[x+y*size] = cur
		cur++
		if cur == size*size {
			break
		}
		if x+ix == size || x+ix < 0 || (ix != 0 && goal[x+ix+y*size] != 0) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && goal[x+(y+iy)*size] != 0) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy
	}
	return goal
}
