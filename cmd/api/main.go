package main

import (
	"fmt"
	"net/http"
	"timetracker/internal/auth"
	"timetracker/internal/config"
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
	// database_service connection
	dbService := db.NewService(config.Config.DSN)
	dbConn := dbService.GetDB()

	sessionStore := auth.NewFileSystemStore(auth.SessionOptions{
		CookiesKey: config.Config.CookiesAuthSecret,
		MaxAge:     config.Config.CookiesAuthAgeInSeconds,
		Secure:     config.Config.CookiesAuthIsSecure,
		HttpOnly:   config.Config.CookiesAuthIsHttpOnly,
	})

	auth.NewAuthService(sessionStore)

	// queries instance from sqlc
	queries := db.New(dbConn)

	// Initialize repositories
	projectRepo := repository.NewProjectRepository(dbConn, queries)
	timeEntryRepo := repository.NewTimeEntryRepository(dbConn, queries)

	// services
	projectService := service.NewProjectService(projectRepo)
	timeEntryService := service.NewTimeEntryService(timeEntryRepo)

	// Initialize a validator instance
	var validate = validator.New()

	// handlers
	projectHandler := handler.NewProjectHandler(projectService, validate)
	timeEntryHandler := handler.NewTimeEntryHandler(timeEntryService)

	// router
	r := chi.NewRouter()

	// Middleware for logging
	r.Use(middleware.Logger)

	// Static files (no auth required)
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Public routes (no authentication required)
	r.Group(func(public chi.Router) {
		public.Get("/auth/login", handler.Make(handler.HandleLoginPage))
		public.Get("/auth/{provider}", handler.Make(handler.HandleProviderLogin))
		public.Get("/auth/{provider}/callback", handler.Make(handler.HandleAuthCallbackFunction))
		public.Get("/auth/logout", handler.Make(handler.HandleLogout))
	})

	// Protected routes (authentication required)
	r.Group(func(protected chi.Router) {
		protected.Use(auth.RequireAuth) // Apply RequireAuth middleware to these routes

		// Home and Projects
		protected.Get("/", handler.Make(handler.RenderHomeIndex))
		protected.Get("/projects", handler.Make(projectHandler.RenderProjectList))
		protected.Get("/projects/new", handler.Make(projectHandler.RenderProjectForm))
		protected.Post("/projects", handler.Make(projectHandler.HandleProjectSubmit))

		// Time Entries
		protected.Get("/timer", handler.Make(timeEntryHandler.RenderTimeEntryIndex))
	})

	// start server
	server := server.NewServer(r, config.Config.Port)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
