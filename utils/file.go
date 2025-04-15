package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func MultipartFileSave(path string, fileName string, file multipart.File) error {
	imageData, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	imagePath := path + fileName
	err = os.WriteFile(imagePath, imageData, 0644)
	if err != nil {
		return err
	}
	return nil
}
