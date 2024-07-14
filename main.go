package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", showList)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func showList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
