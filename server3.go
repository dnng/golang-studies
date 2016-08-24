// gopl.io/ch1/server2
// server2 is minimal "echo" and counter server
package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"sync"
	"strconv"
	"./lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		cyc, err:= strconv.Atoi(r.FormValue("cycles" ))
		if err != nil {
			fmt.Println(cyc)
			fmt.Println(err)
			fmt.Fprintf(os.Stderr, "server3: %v\n", err)
		}
		fmt.Println("start")
		fmt.Println(cyc)
		fmt.Println(cyc)
		fmt.Println(cyc)
		fmt.Println("end")
		lissajous.Lissajous(w, cyc)
	})
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echose the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
