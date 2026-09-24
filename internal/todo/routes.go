package todo

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, pref string, db *sql.DB) {
	store := NewStore(db)
	todos := NewTodos(store)
	handler := NewHandler(todos)

	routes := handler.Routes()
	mux.Handle(pref+"/todo/", http.StripPrefix(pref+"/todo", routes))
}
