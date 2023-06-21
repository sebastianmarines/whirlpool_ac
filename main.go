package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Path("/").HandlerFunc(HomeHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}
