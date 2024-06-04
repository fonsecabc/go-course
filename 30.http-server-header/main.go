package main

import "net/http"

type Controller func(writer http.ResponseWriter, request *http.Request)

type Route struct {
	Path       string
	Method     string
	Controller Controller
}

var routes []Route = []Route{
	{
		Path:       "/",
		Method:     "GET",
		Controller: HelloWorldController,
	},
}

const PORT = ":8080"

func main() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Controller)
	}

	http.ListenAndServe(PORT, nil)
}

func HelloWorldController(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("404 - Not Found"))
		return
	}

	writer.Write([]byte("Hello, World!"))
	writer.WriteHeader(http.StatusOK)
}
