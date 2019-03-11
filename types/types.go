package types

import (
	"crypto/rsa"
	"math/big"
)

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
	NeighborKey    rsa.PublicKey
	ActiveProposal string
	ActiveContract string
	Proposals      []string
	Contracts      []string
}

type Proposal struct {
	From  string
	To    string
	Price float64
	Salt  [4]byte
	TTL   uint64 ///seconds
}

type Contract struct {
	LastSertificate struct {
		Amount    big.Int
		Signature string
	}
	EthAddress  string
	TotalAmount big.Int
	Socket      string
}

type UserData struct {
	username     string
	passwordHash [32]byte /// SHA25
	Sockets      []string
	PublicKey    rsa.PublicKey
	PrivateKey   rsa.PrivateKey
}
