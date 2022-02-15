package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprintln(writer, "Hello World")
		})
	_ = http.ListenAndServe("127.0.0.1:8090", nil)
}
