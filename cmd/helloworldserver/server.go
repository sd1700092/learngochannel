package main

import (
	"net/http"
	"fmt"
)

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "Hello World %s\n", request.FormValue("name"))
	//	return
	//})
	//http.ListenAndServe(":8888", nil)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world, %s\n", request.FormValue("name"))
	})
	http.ListenAndServe(":8888", nil)
}
