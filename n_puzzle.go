package main
import (
	"bufio"
	"fmt" // for print
	"os" // for Args and exit
	"strconv"
	"strings"
	"errors"
	//"n_puzzle/algo"	
)
func checkFormat(data []byte) bool {
	fmt.Printf("string read:\n%s\n", data)
	/*
	split by \n
	check we only need numbers # \t spaces and \n
	kick lines beginning with #
	forget string between # and \n
	check if we have a line with size at the very beginning
	check if we only have numbers  before \n or #, and we should have x numbers by line and x lines of numbers (x = size)
	if all those lines are good we have a good format, we can try to solve this puzzle
	*/
	return true
}

func readFile(name string) ([]int, int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, 0,  errors.New("No such file or directory")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for strings.Trim(scanner.Text(), "\n \t")[0] == '#' {
		scanner.Scan()
	}
	// should we manage "correct maps" without size defined?
	size, err := strconv.Atoi(strings.Trim(scanner.Text(), "\n \t"))
	if err != nil {
		return nil, 0, err
	}
	if size < 2 {
		return nil, 0, errors.New("Puzzle size must be at least 2")
	}
	initialState := make([]int, size*size) 
	// Should we represent it as an array of int or an array of array of int already there?
	// Both options are possible, for the first one we can play with %size to get the tile upside and downside
	// for the second one is [i][j] -1 or +1
	// i guess we can manage it with only one array good challenge
	i := 0
	for scanner.Scan() {
		if strings.Trim(scanner.Text(), "\n \t")[0] == '#' {
			continue
		}
		if i == size {
			return nil, 0,  errors.New("Too many rows")
		}

		parts := strings.Split(strings.Trim(scanner.Text(), "\n \t"), " ")
		if len(parts) < size {
			return nil, 0, errors.New("Row too short")
		}
		if (len(parts) > size && strings.Trim(parts[size], "\n \t")[0] != '#') {
			return nil, 0, errors.New("Row too long")
		}

		for j := 0; j < size; j++ {
			initialState[i*size+j], err = strconv.Atoi(parts[j])
			if err != nil {
				return nil, 0,  err
			}
		}
		i++
	}
	if i < size {
		return nil, 0, errors.New("Not enough rows")
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	return initialState, size, nil
}

// by default the goal is the snail pattern, we could manage differents ways as an option
// by default the puzzle goes from 1 to size * size -1
// this is an attempt of all in one array, we could do it with [][]

func createSolvedState(size int) []int{
	totalNbrTiles := size * size
	array := make([]int, totalNbrTiles)
	value := 1
	nbrTurns := 0
	for thereIsStillAZero(array) {
		for i := size * nbrTurns; i < size * (nbrTurns + 1); i++ {
			if array [i] == 0 {
				array[i] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
//		fmt.Println("apres 1ere ligne")
		for j := size * nbrTurns - 1; j < len(array); j++ {
			if j % size == size - nbrTurns - 1 && array[j] == 0 {
				array[j] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
//		fmt.Println("apres last column")
		for k := len(array) - nbrTurns * size - 1; k >= len(array) - size * (nbrTurns + 1) ; k-- {
			if array[k] == 0 {
				array[k] = value
				value++
			}
		}
		// printSquareFromArray(array, size)
//		fmt.Println("apres last ligne")
		for l := totalNbrTiles - nbrTurns * size - 1; l > 0; l-- {
			if l % size == nbrTurns && array[l] == 0 {
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

// Stop condition to the create solved state function
func thereIsStillAZero(array []int) bool {
	for i:=0; i < len(array); i++ {
		if array[i] == 0 {
			
//		fmt.Println("there is still a zero at index",i)
			return true
		}
	}
//	fmt.Println("there is no more zero")
	return false
}

// print square from array takes an array of int and a size and diplay it
func printSquareFromArray(array []int, size int){
	for i:=1; i <= size; i++ {
		fmt.Println(array[(size * (i - 1)) :(size * i)])
	}
}

func main() {
	if 2 == len(os.Args) {
		input, size, err := readFile(os.Args[1])
		if err != nil {
			fmt.Printf("\033[1;31m%s\033[m\n", err)
		} else {
			fmt.Println(input);
			goalState := createSolvedState(size)
			currentState := input
			fmt.Println("goalState")
			printSquareFromArray(goalState, size)
			printSquareFromArray(currentState, size)
			sigmaDistance, err := CalculSigmaDistance(goalState, currentState, size)
			fmt.Printf("sigmaDistance here :%d\n and err %v\n", sigmaDistance, err)
		}
	} else {
		fmt.Printf("\033[1;31mPlease put only one file in argument, currently, there is %d argument(s)\033[m\n", len(os.Args)-1)
	}
	
	return
	/*
	handle args
	handle options (put into global env)
	check if file is well formated
	check if file is solvable
	solve and display solutions
	return 😉
	*/
}


func CalculSigmaDistance(solvedArray []int, currentArray []int, size int) (sigmaDistance int, err error) {
	sigmaDistance = 0
	for i := 0; i < len(currentArray) - 1; i++ {
		value := solvedArray[i]
		currentPosition, err := returnIndexOfAValue(value, currentArray)
		if err != nil {
			sigmaDistance += getDistanceBetweenTwoPoints(i, currentPosition, size)
		} else {
			return sigmaDistance, err
		}
	}
	return sigmaDistance, err
}

func returnIndexOfAValue(value int, currentArray []int) (result int, err error) {
	for i := 0; i < len(currentArray) - 1; i++ {
		if currentArray[i] == value {
			result = i
			return result, nil
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

func absolute(negativeNbr int) int {
	positiveNumber := negativeNbr
	if negativeNbr < 0 {
		positiveNumber = positiveNumber * -1
	}
	return positiveNumber
}