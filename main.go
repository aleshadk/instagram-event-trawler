package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT");
	if (port != "") {
		return port
	}

	return "8080"
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! " + time.Now().String())
}

func main() {
	http.HandleFunc("/", handler)
	port := getPort()
	fmt.Println("Server started at http://localhost:" + port)
	http.ListenAndServe(":" + port, nil)
}
