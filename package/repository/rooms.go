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
	query := fmt.Sprintf("INSERT INTO %s (id, shortId, roundDuration, roomDuration, roomStartDate, roundStartDate) VALUES ($1, $2, $3, $4, $5, $6)", Table_Rooms)
	if _, err := s.db.Exec(query, input.Wallet, input.Wallet[len(input.Wallet)-6:], 1*60*1000, 15*60*1000, time.Now().Unix(), 0); err != nil {
		return data, err
	}
	data.Id = input.Wallet
	data.ShortId = input.Wallet[len(input.Wallet)-6:]
	data.RoomDuration = 15 * 60 * 1000
	data.RoundDuration = 5 * 60 * 1000
	data.RoomStartDate = time.Now().Unix()
	data.RoundStartDate = 0
	return data, nil
}

func (s RoomsPostgres) DeleteRoom(wallet_id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", Table_Rooms)
	if _, err := s.db.Exec(query, wallet_id); err != nil {
		return err
	}
	return nil
}
func (s RoomsPostgres) AddUserToRoom(input shotball.JoinData) (shotball.CreatedRoom, error) {

	var data shotball.CreatedRoom
	query := fmt.Sprintf("SELECT * FROM %s WHERE shortId=$1", Table_Rooms)
	err := s.db.Get(&data, query, input.Room_id)
	if err != nil {
		return data, err
	}
	query1 := fmt.Sprintf("INSERT INTO %s (id_room, id_user, id_wallet, score) VALUES ($1,$2,$3,0)", Table_users)
	_, err = s.db.Exec(query1, data.Id, input.Telegram, input.Wallet)
	if err != nil {
		return data, nil
	}
	query2 := fmt.Sprint("UPDATE %s SET roundStartDate=$1 WHERE id_room=$2", Table_Rooms)
	_, err = s.db.Exec(query2, time.Now().Unix(), data.Id)
	if err != nil {
		return data, nil
	}

	return data, nil
}
func (s RoomsPostgres) UpdateRoom(data shotball.IncrementData) error {
	query := fmt.Sprintf("UPDATE %s SET score=score+1 WHERE id_room=$1 and id_wallet=$2", Table_users)

	_, err := s.db.Exec(query, data.Room_id, data.Wallet)
	if err != nil {
		return err
	}

	return err
}
