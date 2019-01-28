package main

import (
	"log"

	// my api all resources import.
	api "./scope"
)

func main() {
	log.Println("API Server Started!!")

	router := api.GinRouter()

	router.Run(":8080")

	//log.Fatal(http.ListenAndServe(":8080", router))
}
