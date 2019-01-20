package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (st *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := st.Greeting + st.Punct + st.Who
	fmt.Fprint(w, res)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gopher"})
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("run Go Server run!")
}
