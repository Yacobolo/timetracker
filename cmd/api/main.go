package main

import (
	"fmt"
	"net/http"
	"timetracker/internal/db"
	"timetracker/internal/handler"
	"timetracker/internal/repository"
	"timetracker/internal/server"
	"timetracker/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

func main() {
	// database connection
	database := db.NewService()

	queries := db.New(database.GetDB())

	// Initialize repositories

	projectRepo := repository.NewProjectRepository(queries)

	// services
	projectService := service.NewProjectService(projectRepo)

	// Initialize a validator instance
	var validate = validator.New()

	// handlers
	projectHandler := handler.NewProjectHandler(projectService, validate)

	// router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/", handler.Make(handler.RenderHomeIndex))
	r.Get("/projects", handler.Make(projectHandler.RenderProjectList))
	r.Post("/projects", handler.Make(projectHandler.HandleProjectSubmit))
	r.Get("/project-form", handler.Make(projectHandler.RenderProjectForm))

	// start server
	server := server.NewServer(database, r)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
