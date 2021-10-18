package main

import (
	"apalabrados/controllers"
	"apalabrados/models"
	"log"
	"net/http"
)

func main() {
	models.ConnectDB()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	controllers.RegisterControllers()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
