package todo

import (
	"errors"
	"learn/internal/utils"
	"log"
	"net/http"
	"uuid"
)

type Handler struct {
	Todos *TODOs
}

func NewHandler(todos *TODOs) *Handler {
	return &Handler{todos}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{id}", h.HandleGetTodo)
	mux.HandleFunc("GET /all", h.HandleGetAllTodo)
	mux.HandleFunc("POST /add", h.HandleAddTodo)
	mux.HandleFunc("DELETE /{id}", h.HandleDeleteTodo)
	mux.HandleFunc("PUT /{id}", h.HandleUpdateTodo)
	return mux
}

func (h *Handler) HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{
			"error": ErrInvalidID.Error(),
		})
		return
	}

	todo, err := h.Todos.Get(id)
	if errors.Is(err, ErrTodoNotFound) {
		log.Print(err)
		utils.WriteJson(w, http.StatusNotFound, utils.Envelope{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{
			"error": "failed to get todo",
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, todo)
}

func (h *Handler) HandleGetAllTodo(w http.ResponseWriter, r *http.Request) {
	all, err := h.Todos.GetAll()
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{
			"error": "failed to fetch all todo",
		})
		return
	}
	utils.WriteJson(w, http.StatusOK, all)
}

func (h *Handler) HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := utils.ReadJson[TODO](r)
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{
			"error": ErrInvalidJson.Error(),
		})
		return
	}

	res, err := h.Todos.Add(todo.Title, todo.Description)
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{
			"error": "failed to add todo",
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, res)
}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{
			"error": ErrInvalidID.Error(),
		})
		return
	}

	if err := h.Todos.Delete(id); err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusNotFound, utils.Envelope{
			"error": ErrTodoNotFound.Error(),
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.Envelope{
		"message": "todo delted succesfully",
	})
}

func (h *Handler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{
			"error": ErrInvalidID.Error(),
		})
		return
	}

	todo, err := utils.ReadJson[TODO](r)
	if err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusBadRequest, utils.Envelope{
			"error": ErrInvalidJson.Error(),
		})
		return
	}
	todo.ID = id

	if err := h.Todos.Update(&todo); err != nil {
		log.Print(err)
		utils.WriteJson(w, http.StatusInternalServerError, utils.Envelope{
			"error": "failed to update todo",
		})
		return
	}

	utils.WriteJson(w, http.StatusOK, utils.Envelope{
		"message": "todo updated succesfully",
	})
}
