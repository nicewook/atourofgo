package main

import (
	"fmt"
	"os"
)

// Reader is Interface
type Reader interface {
	Read(b []byte) (n int, err error)
}

// Writer is Interface
type Writer interface {
	Write(b []byte) (n int, err error)
}

// ReadWriter is Interface
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
