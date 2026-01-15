package main

import (
	"fmt"
	"net/http"

	"github.com/fsiwi-hka/IwIValidAPI/src"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var TOKEN string

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := mux.NewRouter()
	r.Use(src.LoggingMiddleware)
	r.HandleFunc("/GetValidUsers", src.HandlerGetValidUser()).Methods("GET")

	fmt.Println("Starting Server at :8080")
	http.ListenAndServe(":8080", r)
}
