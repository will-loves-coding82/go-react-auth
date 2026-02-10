package auth

import (
	"goAuthExample/pkg/app"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var (
	key    = os.Getenv("SESSION_SECRET")
	maxAge = 3600 // 1 hour
	isProd = false
)

type AuthHandler struct {
	client *app.AuthClient
}

func NewAuth() {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd
	store.Options.SameSite = http.SameSiteLaxMode
	gothic.Store = store

	// Register the Goggle OAuth Provider with a custom
	// callback url that points our go endpoint
	goth.UseProviders(google.New(googleClientId, googleClientSecret, "http://localhost:3000/auth/google/callback"))
}
