package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my web application!</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println(" POST is:: ", port)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, mux)
}
