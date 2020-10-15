package main

import (
	"./algo/resolve"
	"bufio"
	"errors"
	"fmt" // for print
	"os"  // for Args and exit
	"strconv"
	"strings"
	//"n_puzzle/algo"
)

var Size int
var GoalState []int

func readFile(name string) ([]int, int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, 0, errors.New("No such file or directory")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for strings.Trim(scanner.Text(), "\n \t")[0] == '#' {
		scanner.Scan()
	}
	size, err := strconv.Atoi(strings.Trim(scanner.Text(), "\n \t"))
	if err != nil {
		return nil, 0, err
	}
	if size < 2 {
		return nil, 0, errors.New("Puzzle size must be at least 2")
	}
	initialState := make([]int, size*size)
	i := 0
	for scanner.Scan() {
		if strings.Trim(scanner.Text(), "\n \t")[0] == '#' {
			continue
		}
		if i == size {
			return nil, 0, errors.New("Too many rows")
		}

		parts := strings.Split(strings.Trim(scanner.Text(), "\n \t"), " ")
		for idx := 0; idx < len(parts); idx++ {
			if len(parts[i]) == 0 {
				copy(parts[idx:], parts[idx+1:])
				parts = parts[:len(parts)-1]
				idx--
			}
		}
		if len(parts) < size {
			return nil, 0, errors.New("Row too short")
		}
		if len(parts) > size && strings.Trim(parts[size], "\n \t")[0] != '#' {
			return nil, 0, errors.New("Row too long")
		}

		for j := 0; j < size; j++ {
			tmp, err := strconv.Atoi(parts[j])
			if err != nil {
				return nil, 0, err
			}
			if tmp >= size*size || tmp < 0 {
				return nil, 0, errors.New("One of the values is too large")
			}
			initialState[i*size+j] = tmp
		}
		i++
	}
	if i < size {
		return nil, 0, errors.New("Not enough rows")
	}

	for i := 0; i < len(initialState); i++ {
		for j := i + 1; j < len(initialState); j++ {
			if initialState[i] == initialState[j] {
				return nil, 0, errors.New("Duplicate number")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	return initialState, size, nil
}

func main() {
	if 2 == len(os.Args) {
		input, size, err := readFile(os.Args[1])
		if err != nil {
			fmt.Printf("\033[1;31m%s\033[m\n", err)
		} else {
			resolve.Resolve(size, input)
		}
	} else {
		fmt.Printf("\033[1;31mPlease put only one file in argument, currently, there is %d argument(s)\033[m\n", len(os.Args)-1)
	}
	return
}
