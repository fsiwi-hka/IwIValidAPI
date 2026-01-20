package main

import (
	"fmt"
	"net/http"
	"os"

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
	var port string = ":" + os.Getenv("PORT")
	fmt.Printf("Starting Server at %s", port)
	http.ListenAndServe(port, r)
}
