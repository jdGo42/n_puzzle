package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	sizeGlobal3 := 3

	sizeGlobal4 := 4
	//	s2 := sizeGlobal * sizeGlobal
	var arrayInitial3 = []int{2, 7, 8, 5, 6, 4, 3, 1, 0}
	//	var perfectInitial3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arrayInitial4 = []int{2, 7, 8, 5, 6, 4, 3, 1, 0, 9, 10, 11, 13, 12, 14, 15, 16}
	arrayExpected3 := createSolvedState(3)
	arrayExpected4 := createSolvedState(4)
	//	sigmaDistancePerfect3, err := CalculSigmaDistance(perfectInitial3, arrayExpected3, sizeGlobal3)
	sigmaDistance, err := CalculSigmaDistance(arrayInitial3, arrayExpected3, sizeGlobal3)
	fmt.Printf("Hello, \nsigmaDistance :%v,\nerr %v\n", sigmaDistance, err)
	sigmaDistance4, err4 := CalculSigmaDistance(arrayInitial4, arrayExpected4, sizeGlobal4)
	fmt.Printf("Hello, \nsigmaDistance :%v,\nerr2 %v\n", sigmaDistance4, err4)
	return
}

func thereIsStillAZero(array []int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {

			//		fmt.Println("there is still a zero at index",i)
			return true
		}
	}
	//	fmt.Println("there is no more zero")
	return false
}

func CalculSigmaDistance(solvedArray []int, currentArray []int, size int) (sigmaDistance int, err error) {
	sigmaDistance = 0
	err = nil
	for i := 0; i < len(currentArray)-1; i++ {
		value := solvedArray[i]
		//	fmt.Printf("into CalculSigmaDistance here i =%v  value is :%v and err is %v",i,  value, err)
		currentPosition, err := returnIndexOfAValue(value, currentArray)
		if err != nil {
			sigmaDistance += getDistanceBetweenTwoPoints(i, currentPosition, size)
		}
	}
	fmt.Printf("distance between index 2 and index 5 into an array of 3\nExpected : 1\nResult : %d\n", getDistanceBetweenTwoPoints(2, 5, 3))
	fmt.Printf("distance between index 1 and index 4 into an array of 3\nExpected : 1\nResult : %d\n", getDistanceBetweenTwoPoints(1, 4, 3))
	fmt.Printf("distance between index 1 and index 7 into an array of 3\nExpected : 2\nResult : %d\n", getDistanceBetweenTwoPoints(1, 7, 3))
	fmt.Printf("distance between index 0 and index 8 into an array of 3\nExpected : 4\nResult : %d\n", getDistanceBetweenTwoPoints(0, 8, 3))
	fmt.Printf("distance between index 17 and index 19 into an array of 17\nExpected : 2\nResult : %d\n", getDistanceBetweenTwoPoints(17, 19, 17))
	fmt.Printf("distance between index 7 and index 8 into an array of 100\nExpected : 1\nResult : %d\n", getDistanceBetweenTwoPoints(7, 8, 100))
	fmt.Printf("distance between index 0 and index 23 into an array of 7\nExpected : ?\nResult : %d\n", getDistanceBetweenTwoPoints(0, 23, 7))

	return sigmaDistance, err
}

func returnIndexOfAValue(value int, currentArray []int) (result int, err error) {
	for i := 0; i < len(currentArray)-1; i++ {
		if currentArray[i] == value {
			result = i
			err = nil
			return result, err
		}
	}
	return -1, errors.New("Value has not been reached into returnIndexOfAValue")
}

func getDistanceBetweenTwoPoints(index, currentPosition, size int) int {
	distanceInLine := absolute((index % size) - (currentPosition % size))
	distanceInColumn := absolute((index / size) - (currentPosition / size))
	totalDistance := distanceInColumn + distanceInLine
	return totalDistance
}

func absolute(n int) int {
	y := n >> (strconv.IntSize - 1)
	return (n ^ y) - y
}

// func randomState(size int) []int {
// 	for i:= 0; i < (size * size); i++ {
// 		var randomNumbers = []int
// 	}
// }

func createSolvedState(size int) []int {
	totalNbrTiles := size * size
	array := make([]int, totalNbrTiles)
	value := 1
	nbrTurns := 0
	for thereIsStillAZero(array) {
		for i := size * nbrTurns; i < size*(nbrTurns+1); i++ {
			if array[i] == 0 {
				array[i] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
		//		fmt.Println("apres 1ere ligne")
		for j := size*nbrTurns - 1; j < len(array); j++ {
			if j%size == size-nbrTurns-1 && array[j] == 0 {
				array[j] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
		//		fmt.Println("apres last column")
		for k := len(array) - nbrTurns*size - 1; k >= len(array)-size*(nbrTurns+1); k-- {
			if array[k] == 0 {
				array[k] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
		//		fmt.Println("apres last ligne")
		for l := totalNbrTiles - nbrTurns*size - 1; l > 0; l-- {
			if l%size == nbrTurns && array[l] == 0 {
				array[l] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
		//		fmt.Println("apres first column")

		nbrTurns++
	}
	// TODO take the maximum and put it to zero to create an empty tile

	return array
}
