package repositoryes

import sqlx "github.com/jmoiron/sqlx"

type Rooms interface {
	// GetAll()
	// GetById()
	// Post()
	// Update()
	// Delete()
}
type Repository struct {
	Rooms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Rooms: NewRoomsPostgres(db),
	}
}
