package services

import (
	"bufio"
	"encoding/csv"
	"os"

	"main/app/models"
	"main/platform/amazon"
)

func S3UploadKey(keyIn *models.UserKey) (string, error) {
	err := createFile(keyIn)
	if err != nil {
		return "", err
	}
	file, err := os.Open(keyIn.Email + ".csv")
	if err != nil {
		return "", err
	}
	msg, err := amazon.S3UploadObject(file)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func S3GetKey(email string) ([]string, error) {
	key, err := amazon.S3GetObject(email + ".csv")
	if err != nil {
		return nil, err
	}
	return key, nil
}

func createFile(keyIn *models.UserKey) error {
	file, err := os.Create(keyIn.Email + ".csv")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(bufio.NewWriter(file))
	err = writer.Write([]string{keyIn.AccessKey, keyIn.SecretKey})
	if err != nil {
		return err
	}
	writer.Flush()
	defer file.Close()
	return nil
}
