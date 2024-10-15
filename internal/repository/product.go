package repository

import "database/sql"

type Repository struct {
	Db *sql.DB
}

type RepositoryInterface interface {
	InsertProduct()
}

func NewRepository() RepositoryInterface {
	return &Repository{
		Db: nil,
	}
}

func (r *Repository) InsertProduct() {
	r.Helper()
	r.Db.Exec("INSERT INTO products (code, price) VALUES ('D42', 100)")
}
func (r *Repository) Helper() {
	r.Db.Exec("INSERT INTO products (code, price) VALUES ('D42', 100)")
}
