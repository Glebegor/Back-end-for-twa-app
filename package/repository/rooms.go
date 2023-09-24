package repositoryes

import (
	"fmt"
	shotball "hackthon_back"
	"time"

	sqlx "github.com/jmoiron/sqlx"
)

type RoomsPostgres struct {
	db *sqlx.DB
}

func NewRoomsPostgres(db *sqlx.DB) *RoomsPostgres {
	return &RoomsPostgres{db: db}
}

func (s RoomsPostgres) CreateRoom(input shotball.CreateData) (shotball.CreatedRoom, error) {
	var data shotball.CreatedRoom
	query := fmt.Sprintf("INSERT INTO %s (id, roundDuration, roomDuration, roomStartDate, roundStartDate) VALUES ($1, $2, $3, $4, $5)", Table_Rooms)
	if _, err := s.db.Exec(query, input.Wallet, 1*60*100, 15*60*100, time.Now().Unix(), time.Now().Unix()); err != nil {
		return shotball.CreatedRoom{}, err
	}
	data.Id = input.Wallet
	data.RoomDuration = 1 * 60 * 100
	data.RoundDuration = 5 * 60 * 100
	data.RoomStartDate = time.Now().Unix()
	data.RoundStartDate = time.Now().Unix()
	return data, nil
}
func (s RoomsPostgres) DeleteRoom(wallet_id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", Table_Rooms)
	if _, err := s.db.Exec(query, wallet_id); err != nil {
		return err
	}
	return nil
}
func (s RoomsPostgres) AddUserToRoom(input shotball.JoinData) (shotball.AllId, error) {
	query := 
	return nil
}
func (s RoomsPostgres) UpdateRoom(data shotball.IncrementData) error {
	query := fmt.Sprintf("UPDATE %s SET score=score+1 WHERE id_room=$1 id_wallet=$2", Table_users)
	if _, err := s.db.Exec(query, data.Room_id, data.Wallet); err != nil {
		return err
	}
	return nil
}
