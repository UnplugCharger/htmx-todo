package handlers

import (
	"github.com/UnplugCharger/htmx-todo/frontend/templates"
	"net/http"
)

type GetHomeHandler struct {
}

func NewGetHomeHandler() *GetHomeHandler {
	return &GetHomeHandler{}
}

// ServeHTTP handles the get home request
func (h *GetHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Index("isaacbyron@gmail.com")

	err := templates.Layout(c, "Home").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
		return
	}
}
