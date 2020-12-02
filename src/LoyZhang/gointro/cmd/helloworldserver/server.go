package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http://localhost:8080/?name=zhanglei
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"<h1>hello word %s</h1>", request.FormValue("name"))
	})
	http.ListenAndServe(":8080",nil)

}
