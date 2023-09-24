package repositoryes

import (
	shotball "hackthon_back"

	sqlx "github.com/jmoiron/sqlx"
)

type Rooms interface {
	CreateRoom(input shotball.CreateData) (shotball.CreatedRoom, error)
	DeleteRoom(wallet_id string) error
}
type Repository struct {
	Rooms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Rooms: NewRoomsPostgres(db),
	}
}
