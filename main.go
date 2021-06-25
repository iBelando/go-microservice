package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iBelando/go-microservice/commons"
	"github.com/iBelando/go-microservice/routes"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 8000")
	log.Println(server.ListenAndServe())
}
