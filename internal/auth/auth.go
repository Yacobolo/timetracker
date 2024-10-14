// package auth

// import (
// 	"os"

// 	"github.com/gorilla/sessions"
// 	"github.com/markbates/goth"
// 	"github.com/markbates/goth/gothic"
// 	"github.com/markbates/goth/providers/azuread"
// )

// const (
// 	key    = "secret"
// 	MaxAge = 300
// 	IsProd = false
// )

// func NewAuth() {
// 	azureClientId := os.Getenv("AZURE_CLIENT_ID")
// 	azureClientSecret := os.Getenv("AZURE_CLIENT_SECRET")

// 	store := sessions.NewCookieStore([]byte(key))
// 	store.MaxAge(MaxAge)

// 	store.Options.Path = "/"
// 	store.Options.Secure = IsProd
// 	store.Options.HttpOnly = true

// 	gothic.Store = store

// 	goth.UseProviders(
// 		azuread.New(azureClientId, azureClientSecret, "http://localhost:3000/auth/azuread/callback", nil),
// 	)
// }

package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"timetracker/internal/config"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
)

func NewAuthService(store sessions.Store) {
	gothic.Store = store

	goth.UseProviders(
		azureadv2.New(
			config.Config.AzureADClientID,
			config.Config.AzureADClientSecret,
			buildCallbackURL("azureadv2"),
			azureadv2.ProviderOptions{
				Tenant: azureadv2.TenantType(config.Config.AzureADTenantID),
				Scopes: []azureadv2.ScopeType{
					"User.Read",
				},
			},
		),
	)
}

func GetSessionUser(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, err
	}

	u := session.Values["user"]
	if u == nil {
		return goth.User{}, fmt.Errorf("user is not authenticated! %v", u)
	}

	return u.(goth.User), nil
}

func StoreUserSession(w http.ResponseWriter, r *http.Request, u goth.User) error {
	session, _ := gothic.Store.Get(r, SessionName)

	session.Values["user"] = u
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func RemoveUserSession(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = goth.User{}
	session.Options.MaxAge = -1

	session.Save(r, w)
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSessionUser(r)
		if err != nil {
			log.Println("User not authenticated:", err)
			http.Redirect(w, r, "/auth/login", http.StatusTemporaryRedirect)
			return
		}

		log.Printf("User is authenticated! user: %v!", session.FirstName)
		next.ServeHTTP(w, r)
	})
}

func buildCallbackURL(provider string) string {
	return fmt.Sprintf("%s:%s/auth/%s/callback", config.Config.PublicHost, strconv.Itoa(config.Config.Port), provider)
}
