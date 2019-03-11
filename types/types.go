package types

type SensorPacket struct {
	SensorID  string `json:"sensorID"`
	Total     uint64 `json:"total"`
	Timestamp int64  ///nanoseconds
}

//======== BACKEND ONLY ============

type SocketInfo struct {
	SensorID     string
	Owner        string
	Alias        string
	NeighborAddr string
	NeighborKey  string
}

type UserData struct {
	username     string
	passwordHash [32]byte /// SHA256
}
