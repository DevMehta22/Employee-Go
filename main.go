package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DevMehta22/mongoapi/routes"
)

func main()  {
	fmt.Println("Starting Server...")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
}