// gopl.io/ch3/basename1
// basename removes directory components and a .suffix
// e.g, a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s :="a/b/c.go"
	s1 :="c.d.go"
	s2 :="abc"
	fmt.Println("Testing basename function (by hand)")
	s3 := basename(s)
	s4 := basename(s1)
	s5 := basename(s2)
	fmt.Println(s3, s4, s5)

	fmt.Println("Testing basename function using strings pkg")
	s6 := basename(s)
	s7 := basename(s1)
	s8 := basename(s2)
	fmt.Println(s6, s7, s8)
}
