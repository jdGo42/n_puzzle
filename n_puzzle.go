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
			return nil, 0, errors.New("Too many rows")
		}

		parts := strings.Split(strings.Trim(scanner.Text(), "\n \t"), " ")
		if len(parts) < size {
			return nil, 0, errors.New("Row too short")
		}
		if len(parts) > size && strings.Trim(parts[size], "\n \t")[0] != '#' {
			return nil, 0, errors.New("Row too long")
		}

		for j := 0; j < size; j++ {
			initialState[i*size+j], err = strconv.Atoi(parts[j])
			if err != nil {
				return nil, 0, err
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
