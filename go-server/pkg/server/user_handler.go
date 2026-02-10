package server

import (
	"goAuthExample/pkg/app"
	"goAuthExample/pkg/responses"
	"net/http"

	"github.com/go-chi/render"
)

type UserHandler struct {
	client *app.UserClient
}

func NewUserHandler(client *app.UserClient) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)
	user, getUserByIdErr := h.client.GetUserById(userId)
	if getUserByIdErr != nil {
		render.Render(w, r, responses.ErrServerError(getUserByIdErr))
		return
	}

	render.Render(w, r, responses.UserResponseRender(user))
}
