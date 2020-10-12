package main
import (
	"bufio"
	"fmt" // for print
	"log"
	"os" // for Args and exit
	"strconv"
	"strings"
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
func main() {
	if 2 == len(os.Args) {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Print("\n\033[1;31mNo such file or directory\033[m\n")
			log.Fatal(err)
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
			log.Fatal(err)
		}
		initial_state := make([]int, size*size)
		i := 0
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			if strings.Trim(scanner.Text(), " \t")[0] == '#' {
				continue
			}
			if i == size {
				log.Fatal("Wrong number of row")
			}
			parts := strings.Split(strings.Trim(scanner.Text(), "\n \t"), " ")
			if len(parts) < size || (len(parts) > size && strings.Trim(parts[size], " \t")[0] != '#') {
				log.Fatal("Wrong row size")
			}
			for j := 0; j < size; j++ {
				initial_state[i*size+j], err = strconv.Atoi(parts[j])
				if err != nil {
					log.Fatal(err)
				}
			}
			i++
		}
		if i < size {
			log.Fatal("Wrong number of row")
		}
		fmt.Println(initial_state)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("\n\033[1;31mPlease put only one file in argument, currently, there is %d argument(s)\033[m\n", len(os.Args)-1)
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
