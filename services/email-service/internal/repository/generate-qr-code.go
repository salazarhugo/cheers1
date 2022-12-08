package repository

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func GenerateQRCodePng(data string) ([]byte, error) {
	bytes, err := qrcode.Encode(data, qrcode.High, 256)

	if err != nil {
		log.Fatalln("Error generating QR code. ", err)
		return nil, err
	}

	return bytes, nil
}
