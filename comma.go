// gopl.io/ch3/comma
package main

import (
	"fmt"
	"strconv"
)

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}


func main() {
	dec := 123456
	fmt.Println(comma(strconv.Itoa(dec)))
}

