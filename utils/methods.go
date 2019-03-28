package utils

import (
	"strconv"
	"strings"

	eth "github.com/ethereum/go-ethereum/crypto"
)

func ConvertIPToBytes(IP string) (res []byte) {
	octets := strings.FieldsFunc(IP, IPSplitter)
	for _, octet := range octets {
		oct, err := strconv.Atoi(octet)
		CheckFatal(err)
		res = append(res, byte(oct))
	}
	return
}

func IPSplitter(r rune) bool {
	return r == '.'
}

func PrivateHexToAddress(private string) string {
	priv, err := eth.HexToECDSA(private)
	if err != nil {
		return ""
	}
	return eth.PubkeyToAddress(priv.PublicKey).Hex()
}
