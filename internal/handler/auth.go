package handler

import (
	"context"
	"log"
	"net/http"
	"timetracker/internal/auth"
	"timetracker/internal/templates/pages"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) error {
	_, err := auth.GetSessionUser(r)
	if err != nil {
		log.Println(err)
		return pages.LoginPage().Render(r.Context(), w)
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}

func HandleProviderLogin(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")
	if provider == "" {
		http.Error(w, "Provider not found", http.StatusBadRequest)
		return nil
	}

	// Add the provider to the context for Gothic to use
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Printf("User already authenticated! %v", u)
		pages.LoginPage().Render(r.Context(), w)
		return nil
	} else {
		gothic.BeginAuthHandler(w, r)
		return nil
	}
}

func HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) error {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	if err := auth.StoreUserSession(w, r, user); err != nil {
		return err
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}

func HandleLogout(w http.ResponseWriter, r *http.Request) error {
	log.Println("Logging out...")

	if err := gothic.Logout(w, r); err != nil {
		log.Println(err)
	}

	auth.RemoveUserSession(w, r)

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}
