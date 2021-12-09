package main

import (
	"net/http"
	"todo/config"
	"todo/controller"
	"todo/repository"
	"todo/service"

	"github.com/rs/cors"
)

func main() {
	configuration := config.New()
	db := config.NewDb(configuration)
	mux := http.NewServeMux()

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)
	todoController.Router(mux)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
	})

	http.ListenAndServe(":4000", c.Handler(mux))
}
