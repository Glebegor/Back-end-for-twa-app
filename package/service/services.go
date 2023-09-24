package services

import (
	shotball "hackthon_back"
	repositoryes "hackthon_back/package/repository"
)

type Rooms interface {
	CreateRoom(data shotball.CreateData) (shotball.CreatedRoom, error) // id string, roundDuration 5min, roomDuration 15min, roomStartDate ms, roundStartDate ms
	JoinRoom(data shotball.JoinData) (int32, error)                    // id_room
	IncrementingRoom(data shotball.IncrementData) error
}
type Service struct {
	Rooms
}

func NewService(repo *repositoryes.Repository) *Service {
	return &Service{
		Rooms: NewRoomsService(repo.Rooms),
	}
}
