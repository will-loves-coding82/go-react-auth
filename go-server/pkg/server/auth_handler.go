package server

import (
	"context"
	"goAuthExample/pkg/app"
	"goAuthExample/pkg/responses"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	client *app.AuthClient
}

func NewAuthHandler(client *app.AuthClient) *AuthHandler {
	return &AuthHandler{
		client: client,
	}
}

// Login will initiate the login process with the IdP
func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	q := r.URL.Query()
	q.Set("prompt", "select_account") // or "consent" if you want re-consent
	r.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(w, r)
}

// Callback will be invoked once the user signs in with the IdP
// in order to retrieve their identity metadata.
func (a *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	gothUser, completeAuthErr := gothic.CompleteUserAuth(w, r)

	if completeAuthErr != nil {
		log.Printf(completeAuthErr.Error())
		http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
		return
	}

	session, getSessionErr := gothic.Store.Get(r, "session")
	if getSessionErr != nil {
		log.Printf("Error making new session: %v", getSessionErr)
		http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
		return
	}

	// Get the user profile if possible to perform a login.
	// Otherwise provision a new account in the database
	log.Printf("Getting or creating user")
	userId, getOrCreateUserErr := a.client.GetOrCreateUser(gothUser)
	if getOrCreateUserErr != nil {
		http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
		return
	}

	// Save the user's id in the browser cookie
	session.Values["user_id"] = userId
	saveSessionErr := session.Save(r, w)
	if saveSessionErr != nil {
		log.Printf("Error saving new session: %v", saveSessionErr)
		http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
		return
	}

	// Save the user details for the associated email
	log.Printf("Saving user details")
	updateUserErr := a.client.UpdateUserDetails(gothUser.UserID, gothUser.Email, userId)
	if updateUserErr != nil {
		http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
		return
	}

	http.Redirect(w, r, "http://localhost:8080/dashboard", http.StatusFound)
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, getSessionErr := gothic.Store.Get(r, "session")
	if getSessionErr != nil {
		log.Printf("Error getting session for logout: %v", getSessionErr)
		http.Error(w, "Logout error", http.StatusInternalServerError)
		return
	}

	gothicLogoutErr := gothic.Logout(w, r)
	if gothicLogoutErr != nil {
		log.Printf("Error invalidating gothic session: %v", gothicLogoutErr)
		http.Error(w, "Logout error", http.StatusInternalServerError)
		return
	}
	// Destroy the session by setting its MaxAge to a negative value
	// and clearing its values. This also tells the browser to delete the cookie.
	session.Options.MaxAge = -1
	session.Values = make(map[any]any)   // Clear stored values
	saveSessionErr := session.Save(r, w) // Save the modified session to apply changes (delete cookie)

	if saveSessionErr != nil {
		log.Printf("Error saving session during logout: %v", saveSessionErr)
		http.Error(w, "Logout error", http.StatusInternalServerError)
		return
	}

	log.Printf("Logout successful")
	render.Render(w, r, responses.PostResponseRender("Logged out successfully"))
}
