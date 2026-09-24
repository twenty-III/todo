package todo

import "uuid"

type TODO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

type Storage interface {
	Get(id uuid.UUID) (*TODO, error)
	GetAll() ([]TODO, error)
	Create(t *TODO) (*TODO, error)
	Delete(id uuid.UUID) error
	Update(t *TODO) error
}

type TODOs struct {
	store Storage
}

func NewTodos(store Storage) *TODOs {
	return &TODOs{store}
}

func (ts *TODOs) Get(id uuid.UUID) (*TODO, error) {
	return ts.store.Get(id)
}

func (ts *TODOs) GetAll() ([]TODO, error) {
	return ts.store.GetAll()
}

func (ts *TODOs) Add(title, description string) (*TODO, error) {
	todo := TODO{
		Title:       title,
		Description: description,
		Done:        false,
	}
	return ts.store.Create(&todo)
}

func (ts *TODOs) Delete(id uuid.UUID) error {
	return ts.store.Delete(id)
}

func (ts *TODOs) Update(t *TODO) error {
	return ts.store.Update(t)
}
