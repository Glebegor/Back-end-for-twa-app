package shotball

type CreateData struct {
	Wallet   string `json:"wallet"`
	Telegram string `json:"telegram"`
}
type CreatedRoom struct {
	Id             string `json:"id"`
	RoundDuration  int32  `json:"roundDuration"`  // ms
	RoomDuration   int32  `json:"roomDuration"`   // ms
	RoomStartDate  int64  `json:"roomStartDate"`  // ms
	RoundStartDate int64  `json:"roundStartDate"` // ms
}

type JoinData struct {
	Wallet   string `json:"wallet"`
	Telegram string `json:"telegram"`
	Room_id  string `json: "room_id"`
}
type IncrementData struct {
	Wallet   string `json:"wallet"`
	Telegram string `json:"telegram"`
	Room_id  string `json: "room_id"`
}
