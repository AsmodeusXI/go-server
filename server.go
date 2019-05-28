package main

import (
	"fmt"
	"net/http"
	"go-server/config"
)

func home(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Welcome home!"))
}

func ping(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("ping!"))
}

func main() {
	config, err := config.LoadConfig()
	if err != nil { panic(err) }

	http.HandleFunc("/", home)
	http.HandleFunc("/ping", ping)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		panic(err)
	} else {
		fmt.Printf("Server started!")
	}
}
