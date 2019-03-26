package utils

import (
	"strconv"
	"strings"
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