package programador

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert() (*sql.Stmt, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO programadores (id, apelido, nome, nascimento, stack)
		VALUES ($1, $2, $3, $4, $5)
	`)
	return stmt, err
}

func (r *Repository) Count() (int, error) {
	row := r.db.QueryRow("SELECT COUNT(*) FROM programadores")
	var count int
	err := row.Scan(&count)
	return count, err
}
