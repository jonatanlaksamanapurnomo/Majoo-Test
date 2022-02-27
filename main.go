package main

import (
	"log"
	"majoo-be-test/routes"
	services "majoo-be-test/service"
	"majoo-be-test/util"
	"net/http"
)

func main() {

	var db = util.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()
	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
