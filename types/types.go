package types

type SensorPacket struct {
	UserToken     string `json:"user_token"`
	NeighborToken string `json:"neighbor_token"`
	Value         int64  `json:"voltage"`
	Timestamp     int64
}
