package shotball

type CreateData struct {
	Wallet   string `json:"wallet"`
	Telegram string `json:"telegram"`
}
type CreatedRoom struct {
	Id             string
	RoundDuration  int32 // ms
	RoomDuration   int32 // ms
	RoomStartDate  int64 // ms
	RoundStartDate int64 // ms
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
