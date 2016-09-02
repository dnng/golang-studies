package main

import (
	"os"
	"fmt"
	"strconv"
	"popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pc: %v\n", err)
			os.Exit(1)
		}
		var c1 int
		var c2 int
		var c3 int
		c1 = popcount.PopCount(n)
		c2 = popcount.PopCountLoop(n)
		c3 = popcount.PopCountShift(n)
		fmt.Printf("%v, %v, %v\n", c1, c2, c3)
	}
}
