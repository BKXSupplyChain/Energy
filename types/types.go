package types

import (
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
	NeighborKey    string
	ActiveProposal string
	ActiveContract string
	Proposals      []string
	Contracts      []string
}

type Proposal struct {
	ID    string
	Price uint64
	TotalAmount big.Int
	RelError    float64
	AbsError    big.Int
	TTL         uint64 ///seconds
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
	r              [32]byte
	s              [32]byte
	Amount         [32]byte
	ActiveContract string
}

type ProposalMessage struct {
	UserPublicKey      string
	Proposal           Proposal
	NeighbourPublicKey string
}
