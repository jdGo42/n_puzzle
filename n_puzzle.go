package main

import (
	"io/ioutil"	// for open and read files
	"fmt"	// for print
	"os"	// for Args and exit
)

func main() {

	if 2 == len(os.Args) {
    	data, err := ioutil.ReadFile(os.Args[1])
    	if err != nil {
     		fmt.Print("\n\033[1;31mNo such file or directory\033[m\n")
     		os.Exit(-1)
    	}
    	fmt.Printf("string read:%s\n", data)
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