package main

import (
	"fmt"
	"net/http"
)

func home(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Welcome home!"))
}

func ping(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("ping!"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	} else {
		fmt.Printf("Server started!")
	}
}
