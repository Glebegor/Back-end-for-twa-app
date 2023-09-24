package services

import (
	shotball "hackthon_back"
	repositoryes "hackthon_back/package/repository"
)

type Rooms interface {
	CreateRoom(input shotball.CreateData) (shotball.CreatedRoom, error) // id string, roundDuration 5min, roomDuration 15min, roomStartDate ms, roundStartDate ms
	JoinRoom(input shotball.JoinData) (shotball.CreatedRoom, error)     // id_room
	IncrementingRoom(input shotball.IncrementData) error
}
type Service struct {
	Rooms
}

func NewService(repo *repositoryes.Repository) *Service {
	return &Service{
		Rooms: NewRoomsService(repo.Rooms),
	}
}
