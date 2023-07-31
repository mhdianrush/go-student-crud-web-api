package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-student-crud-web-api/config"
	"github.com/mhdianrush/go-student-crud-web-api/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	route := mux.NewRouter()
	router := route.PathPrefix("/api").Subrouter()
	routes.StudentRouter(router)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: router,
	}

	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to database %s", err.Error())
	}

	logger.Printf("server running on port %s", os.Getenv("SERVER_PORT"))
}
