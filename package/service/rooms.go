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
		time.Sleep(3 * time.Second)
		err = s.repo.DeleteRoom(data.Id)
	}()
	return data, err
}
func (s *RoomsService) JoinRoom(data shotball.JoinData) (int32, error) { // id_room
	return 0, nil
}
func (s *RoomsService) IncrementingRoom(data shotball.IncrementData) error {
	return nil
}
