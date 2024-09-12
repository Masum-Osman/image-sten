package image

import (
	"encoding/base64"
)

func BinaryToBase64(binaryData []byte) string {
	return base64.StdEncoding.EncodeToString(binaryData)
}

func Base64ToBinary(base64Text string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Text)
}
