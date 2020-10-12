package main

import (
	"io/ioutil"	// for open and read files
	"fmt"	// for print
	"os"	// for Args and exit
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
    	data, err := ioutil.ReadFile(os.Args[1])
    	if err != nil {
     		fmt.Print("\n\033[1;31mNo such file or directory\033[m\n")
     		os.Exit(-1)
    	}
    	checkFormat(data)
    } else {
    	fmt.Printf("\n\033[1;31mPlease put only one file in argument, currently, there is %d argument(s)\033[m\n", len(os.Args) - 1 )
    }
    return
	/*
	handle args
	handle options (put into global env)
	check if file is well formated
	check if file is solvable
	solve and display solutions
	return ;-)
	*/
}