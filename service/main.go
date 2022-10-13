package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/HugoIr/steve-shop/service/database"
	"github.com/HugoIr/steve-shop/service/server"
	shipperHandler "github.com/HugoIr/steve-shop/service/server/handlers/shipper"
	"github.com/HugoIr/steve-shop/service/shippermodule"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "12345",
		DBName:   "devcamp",
		Port:     5432,
		Host:     "db",
		SSLMode:  "disable",
	}

	// Init DB connection
	log.Println("Initializing DB Connection")
	db := database.GetDatabaseConnection(dbConfig)

	// Init shipper usecase
	log.Println("Initializing Usecase")
	sm := shippermodule.NewProductModule(db)

	// Init shipper handler
	log.Println("Initializing Handler")
	sh := shipperHandler.NewProductHandler(sm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/shipper", sh.AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/shipper/{id}", sh.UpdateProductHandler).Methods(http.MethodPut)
	router.HandleFunc("/shipper/{id}", sh.GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/shipper/{id}", sh.RemoveProductHandler).Methods(http.MethodDelete)
	router.HandleFunc("/shippers", sh.GetProductAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", sh.RootHandler).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022-snd shipper service service is starting...")

	server.Serve(serverConfig, router)
}
