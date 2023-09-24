package services

import (
	shotball "hackthon_back"
	repositoryes "hackthon_back/package/repository"
	"time"
)

type RoomsService struct {
	repo repositoryes.Rooms
}

func NewRoomsService(repo repositoryes.Rooms) *RoomsService {
	return &RoomsService{repo: repo}
}
func (s *RoomsService) CreateRoom(input shotball.CreateData) (shotball.CreatedRoom, error) { // id string, roundDuration 5min, roomDuration 15min, roomStartDate ms, roundStartDate ms
	data, err := s.repo.CreateRoom(input)
	go func() {
		time.Sleep(time.Minute * 1)
		err = s.repo.DeleteRoom(data.Id)
	}()
	return data, err
}
func (s *RoomsService) JoinRoom(input shotball.JoinData) (shotball.AllId, error) { // id_room
	data, err := s.repo.AddUserToRoom(input)
	return data, err
}
func (s *RoomsService) IncrementingRoom(input shotball.IncrementData) error {
	err := s.repo.UpdateRoom(input)
	if err != nil {
		return err
	}
	return err
}
