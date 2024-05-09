package handlers

import (
	"github.com/UnplugCharger/htmx-todo/frontend/templates"
	"net/http"
)

type GetRegisterHandler struct {
}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

// ServeHTTP handles the get register request
func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.RegisterPage()
	err := templates.Layout(c, "Register").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render register page", http.StatusInternalServerError)
		return
	}
}
