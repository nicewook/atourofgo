package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	// ListenAndServe has handler which is interface
	// and that interface has function of ServeHTTP
	http.ListenAndServe("localhost:4000", h)
}
