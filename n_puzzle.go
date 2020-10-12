package main
import (
	"bufio"
	"fmt" // for print
	"os" // for Args and exit
	"strconv"
	"strings"
	"errors"
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

func read_file(name string) ([]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.New("No such file or directory")
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
		return nil, err
	}
	if size < 2 {
		return nil, errors.New("Puzzle size must be at least 2")
	}
	initial_state := make([]int, size*size)
	i := 0
	for scanner.Scan() {
		if strings.Trim(scanner.Text(), "\n \t")[0] == '#' {
			continue
		}
		if i == size {
			return nil, errors.New("Too many rows")
		}

		parts := strings.Split(strings.Trim(scanner.Text(), "\n \t"), " ")
		if len(parts) < size {
			return nil, errors.New("Row too short")
		}
		if (len(parts) > size && strings.Trim(parts[size], "\n \t")[0] != '#') {
			return nil, errors.New("Row too long")
		}

		for j := 0; j < size; j++ {
			initial_state[i*size+j], err = strconv.Atoi(parts[j])
			if err != nil {
				return nil, err
			}
		}
		i++
	}
	if i < size {
		return nil, errors.New("Not enough rows")
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return initial_state, nil
}

func main() {
	if 2 == len(os.Args) {
		input, err := read_file(os.Args[1])
		if err != nil {
			fmt.Printf("\033[1;31m%s\033[m\n", err)
		} else {
			fmt.Println(input);
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
