package handlers

import (
	db "github.com/UnplugCharger/htmx-todo/db/sqlc"
	"github.com/UnplugCharger/htmx-todo/frontend/templates"
	"net/http"
)

type PostLoginHandler struct {
	store db.Store
}

type PostLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewPostLoginHandler(store db.Store) *PostLoginHandler {
	return &PostLoginHandler{store: store}
}

// ServerHTTP handles the post login request

func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := h.store.GetUser(r.Context(), username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		c := templates.LoginError()
		err := c.Render(r.Context(), w)
		if err != nil {
			return
		}
	}
	if user.Password != password {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		err := c.Render(r.Context(), w)
		if err != nil {
			return
		}
	}

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}
