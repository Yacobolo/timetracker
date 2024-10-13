package main

import (
	"fmt"
	"net/http"
	"timetracker/internal/db"
	"timetracker/internal/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// database_service connection
	dbService := db.NewService()

	// dbConn := dbService.GetDB()

	// // queries instance from sqlc
	// queries := db.New(dbConn)

	// Initialize repositories

	// projectRepo := repository.NewProjectRepository(dbConn, queries)
	// timeEntryRepo := repository.NewTimeEntryRepository(dbConn, queries)

	// // services
	// projectService := service.NewProjectService(projectRepo)
	// timeEntryService := service.NewTimeEntryService(timeEntryRepo)

	// // Initialize a validator instance
	// var validate = validator.New()

	// // handlers
	// projectHandler := handler.NewProjectHandler(projectService, validate)
	// timeEntryHandler := handler.NewTimeEntryHandler(timeEntryService)

	// router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// // Projects
	// r.Get("/", handler.Make(handler.RenderHomeIndex))
	// r.Get("/projects", handler.Make(projectHandler.RenderProjectList))
	// r.Get("/projects/new", handler.Make(projectHandler.RenderProjectForm))
	// r.Post("/projects", handler.Make(projectHandler.HandleProjectSubmit))

	// // Time Entries
	// r.Get("/timer", handler.Make(timeEntryHandler.RenderTimeEntryIndex))

	// start server
	server := server.NewServer(dbService, r)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
