
package main

import (
	"fmt"
	"os"
	"strconv"
)

// CracklePop program
func main() {

	ub := 100
	args := os.Args
	if len(args) > 1 {
		// ignore multiple args, just grab first 1
		if len(args) > 2 {
			fmt.Println(".. only the first argument will be read")
		}
		arg := args[1]
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(".. error reading input, using default value")
		} else if n < 0 {
			fmt.Println(".. error: only accepts positive values, using default value")
		} else {
			ub = n
		}
	}

	fmt.Println(".. starting program..\n\nCracklePop from 1..", ub, "\n")


	for i:=1;i<=ub;i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Print("CracklePop\t")
			} else {
				fmt.Print("Crackle\t")
			}
		} else if i % 5 == 0 {
			fmt.Print("Pop\t")
		} else {
			fmt.Print(strconv.Itoa(i), "\t")
		}
	}
	fmt.Println()
}

