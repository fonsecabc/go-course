package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

var visits int = 0

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		visits++
		writer.Write([]byte(fmt.Sprintf("Visits: %d", visits)))
	})
	http.ListenAndServe(PORT, nil)
}
