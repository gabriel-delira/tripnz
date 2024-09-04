package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func checkServerStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "I am alive!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3000", nil)
  	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}