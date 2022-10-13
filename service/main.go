package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/HugoIr/steve-shop/service/database"
	"github.com/HugoIr/steve-shop/service/server"
	shopHandler "github.com/HugoIr/steve-shop/service/server/handlers/shop"
	"github.com/HugoIr/steve-shop/service/shopmodule"
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

	// Init shop usecase
	log.Println("Initializing Usecase")
	sm := shopmodule.NewProductModule(db)

	// Init shop handler
	log.Println("Initializing Handler")
	sh := shopHandler.NewProductHandler(sm)

	router := mux.NewRouter()

	// REST Handlers
	router.HandleFunc("/product", sh.AddProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", sh.UpdateProductHandler).Methods(http.MethodPut)
	router.HandleFunc("/product/{id}", sh.GetProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", sh.RemoveProductHandler).Methods(http.MethodDelete)
	router.HandleFunc("/products", sh.GetProductAllHandler).Methods(http.MethodGet)
	router.HandleFunc("/", sh.RootHandler).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9090,
	}
	log.Println("Devcamp-2022-snd shop service service is starting...")

	server.Serve(serverConfig, router)
}
