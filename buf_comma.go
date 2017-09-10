// gopl.io/ch3/comma
package main

import (
	"fmt"
	"strconv"
	"bytes"
)

// comma inserts commas in a non-negative decimal integer string
func bufComma(s string) string {
	if len(s) <= 3 {
		return s
	}

	var buf bytes.Buffer
	l := len(s)
	init := l % 3

	if init >= 1 {
		buf.WriteString(s[:init])
		buf.WriteString(",")
	}

	for i := init; i < l; i+= 3 {
		buf.WriteString(s[i:i+3])
		if i + 3 < l {
			buf.WriteString(",")
		}
	}

	return buf.String()
}

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
	fmt.Println(bufComma(strconv.Itoa(dec)))
}

