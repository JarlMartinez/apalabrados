package main

import (
	"apalabrados/controllers"
	"apalabrados/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("$PORT not set. Using defautl")
		port = "8080"
	}

	models.ConnectDB()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	controllers.RegisterControllers()

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
