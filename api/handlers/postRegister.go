package handlers

import (
	db "github.com/UnplugCharger/htmx-todo/db/sqlc"
	"net/http"
)

type PostRegisterHandler struct {
	store db.Store
}

type PostRegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewPostRegisterHandler(store db.Store) *PostRegisterHandler {
	return &PostRegisterHandler{store: store}
}

// ServerHTTP handles the post register request
func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	_, err := h.store.CreateUser(r.Context(), db.CreateUserParams{
		Username: username,
		Password: password,
		Email:    email,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusOK)
}
