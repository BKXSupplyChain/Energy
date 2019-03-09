package types

type SensorPacket struct {
	SensorID  string `json:"sensorID"`
	Total     int64  `json:"total"`
	Timestamp int64
}
