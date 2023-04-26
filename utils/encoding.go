package utils

import "encoding/base64"

func Base64Encode(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}
