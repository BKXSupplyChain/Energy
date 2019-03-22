package types

import "math/big"
import "time"

type SensorPacket struct {
	SensorID  string `json:"sensorID"`
	Total     uint64 `json:"total"`
	Timestamp int64  ///nanoseconds
}

//======== BACKEND ONLY ============

type SocketInfo struct {
	SensorID       string
	Owner          string
	Alias          string
	NeighborAddr   string
	NeighborKey    string
	ActiveProposal string
	ActiveContract string
	Proposals      []string
	Contracts      []string
}

type Proposal struct {
	ID    string
	From  string
	To    string
	Price uint64
	TotalAmount big.Int
	RelError uint16
	AbsError big.Int
	TTL   uint64 ///seconds
}

type Contract struct {
	LastSertificate struct {
		Amount    [32]byte
		Signature string
	}
	EthAddress  string
	TotalAmount [32]byte
	Socket      string
}

type UserData struct {
	Username     string   ///ID is crc64(ECMA) of username
	PasswordHash [32]byte /// SHA256
	Sockets      []string
	PrivateKey   string
}

type Certificate struct {
	r	 	*big.Int
	s		*big.Int
	Amount		*big.Int
	ActiveContract	 string
}
