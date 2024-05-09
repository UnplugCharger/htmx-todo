package handlers

import (
	db "github.com/UnplugCharger/htmx-todo/db/sqlc"
	"github.com/UnplugCharger/htmx-todo/frontend/templates"
	"net/http"
)

type GetListTodoHandler struct {
	store db.Store
}

type GetListTodoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func NewGetListTodoHandler(store db.Store) *GetListTodoHandler {
	return &GetListTodoHandler{store: store}
}

// ServerHTTP handles the get list todo request
func (h *GetListTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//limitStr := r.FormValue("limit")
	//offsetStr := r.FormValue("offset")
	//
	////// Convert limit and offset from string to int32
	////limit, err := strconv.ParseInt(limitStr, 10, 32)
	////if err != nil {
	////	w.WriteHeader(http.StatusBadRequest)
	////	return
	////
	////}
	////offset, err := strconv.ParseInt(offsetStr, 10, 32)
	////if err != nil {
	////	w.WriteHeader(http.StatusBadRequest)
	////	return
	////
	////}

	tasks, err := h.store.ListTasks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the list of tasks
	c := templates.ListTodos(tasks)
	err = templates.Layout(c, "List Todos").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render list todos", http.StatusInternalServerError)
		return
	}

}
