package server

import (
	"context"
	"fmt"
	"goAuthExample/pkg/app"
	"goAuthExample/pkg/database"
	"goAuthExample/pkg/responses"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	db := database.New()

	authClient := app.NewAuthClient(db)
	authHandler := NewAuthHandler(authClient)

	userClient := app.NewUserClient(db)
	userHandler := NewUserHandler(userClient)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Publicly available routes
	r.Get("/auth/login", authHandler.Login)
	r.Get("/auth/{provider}/callback", authHandler.Callback)
	r.Post("/auth/logout", authHandler.Logout)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/user", userHandler.GetUserById)
	})

	return r
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, getSessionErr := gothic.Store.Get(r, "session")

		if getSessionErr != nil {
			render.Render(w, r, responses.ErrServerError(getSessionErr))
			http.Redirect(w, r, "http://localhost:8080/errors/500", http.StatusFound)
			return
		}

		if s == nil || s.Values["user_id"] == nil {
			render.Render(w, r, responses.ErrInvalidRequest(fmt.Errorf("No user session")))
			http.Redirect(w, r, "http://localhost:8080/errors/401", http.StatusFound)
			return
		}

		if s.Options.MaxAge < 0 {
			render.Render(w, r, responses.ErrInvalidRequest(fmt.Errorf("Expired session")))
			http.Redirect(w, r, "http://localhost:8080/errors/401", http.StatusFound)
			return
		}

		// Attach user id to request context so that http handlers can
		// process requests efficiently without having to query database
		ctx := context.Background()
		ctx = context.WithValue(ctx, "user_id", s.Values["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
