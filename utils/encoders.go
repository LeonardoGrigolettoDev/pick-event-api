package utils

import "encoding/base64"

func EncodeImageToBase64(image []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(image)
	return encoded, nil
}
