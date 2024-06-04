package main

import "net/http"

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))

	mux.Handle("/", fileServer)
	http.ListenAndServe(port, mux)
}
