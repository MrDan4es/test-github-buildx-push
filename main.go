package main

import (
	"fmt"
	"io"
	"net/http"
)

const addr = ":3333"

func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "12312321 DID IT")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", getHello)

	fmt.Printf("server starting on port %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
