package handlers

import (
	"github.com/UnplugCharger/htmx-todo/frontend/templates"
	"net/http"
)

type GetLoginHandler struct {
}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

// Handle handles the get login request
func (h *GetLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Login("Login")
	err := templates.Layout(c, "Login").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		return
	}
}
