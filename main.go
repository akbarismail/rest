package main

import (
	"fmt"
	"net/http"

	"github.com/akbarismail/rest/client"
	"github.com/akbarismail/rest/server"
)

func main() {
	mux := http.NewServeMux()

	// server
	mux.HandleFunc("/api/student", server.Students)

	// client
	mux.HandleFunc("/page/student", client.Student)

	fmt.Println("starting web server at localhost:8000")
	http.ListenAndServe(":8000", mux)
}
