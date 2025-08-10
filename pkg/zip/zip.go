package zip

import (
	"bytes"
	"crypto/rand"
	"math/big"

	pzip "github.com/alexmullins/zip"
)

type AttachmentData struct {
	Filename string
	Content  []byte
}

func CreatePasswordProtected(attachments []AttachmentData, password string) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := pzip.NewWriter(buf)
	
	for _, attachment := range attachments {
		fileWriter, err := zipWriter.Encrypt(attachment.Filename, password)
		if err != nil {
			return nil, err
		}
		_, err = fileWriter.Write(attachment.Content)
		if err != nil {
			return nil, err
		}
	}
	
	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}
	
	return buf.Bytes(), nil
}

func GeneratePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	
	return string(password)
}
