package todo

import (
	"database/sql"
	"errors"
	"fmt"
	"uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) Get(id uuid.UUID) (*TODO, error) {
	query := `SELECT id, title, description, done FROM todos WHERE id = ?`

	var todo TODO
	err := s.db.QueryRow(query, id.String()).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrTodoNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return &todo, nil
}

func (s *Store) GetAll() ([]TODO, error) {
	query := `SELECT id, title, description, done FROM todos`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all todos: %w", err)
	}
	defer rows.Close()

	todos := make([]TODO, 0)
	for rows.Next() {
		var todo TODO
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done); err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating todos: %w", err)
	}

	return todos, nil
}

func (s *Store) Create(t *TODO) (*TODO, error) {
	t.ID = uuid.New()

	query := `INSERT INTO todos (id, title, description, done) VALUES (?, ?, ?, ?)`

	_, err := s.db.Exec(query, t.ID.String(), t.Title, t.Description, t.Done)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return t, nil
}

func (s *Store) Delete(id uuid.UUID) error {
	query := `DELETE FROM todos WHERE id = ?`

	result, err := s.db.Exec(query, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

func (s *Store) Update(t *TODO) error {
	query := `UPDATE todos SET title = ?, description = ?, done = ? WHERE id = ?`

	result, err := s.db.Exec(query, t.Title, t.Description, t.Done, t.ID.String())
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}
