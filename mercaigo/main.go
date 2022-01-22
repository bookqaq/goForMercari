package mercarigo

import (
	"encoding/base64"
	"encoding/binary"
)

func intToByte(target int) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(target))
	return result
}

func intToBase64URL(target int) string {
	return base64.StdEncoding.EncodeToString(intToByte(target))
}

func stringToBase64URL(target string) string {
	return "1"
}

func dPoPGenerator() {

}
