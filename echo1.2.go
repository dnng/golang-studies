// echo exercise 1.1
package main

import (
    "fmt"
    "os"
)

func main() {
    for i, s := range os.Args {
        fmt.Println(i, s) 
    }
}
