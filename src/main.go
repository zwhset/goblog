package main

import (
	"controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/register", controllers.Register)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
