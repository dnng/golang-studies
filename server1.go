// gopl.io/ch1/server1
// Server 1 is a minimal "echo" server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// if err := r.ParseForm(); err != nil {
	if r.ParseForm() != nil {
		// log.Print(err)
		fmt.Fprintf(w, "AAAAAA")
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
