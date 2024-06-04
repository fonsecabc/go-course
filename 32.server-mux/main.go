package main

import "net/http"

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	http.ListenAndServe(port, mux)
}
