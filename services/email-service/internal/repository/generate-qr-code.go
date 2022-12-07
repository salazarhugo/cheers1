package repository

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
	"log"
)

func GenerateQRCodeHtmlImageTag(data string) string {
	qrCodeImageData, taskError := qrcode.Encode(data, qrcode.High, 256)

	if taskError != nil {
		log.Fatalln("Error generating QR code. ", taskError)
	}

	//qrcode.WriteFile()
	// Encode raw QR code data to base 64
	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)

	return "<img src=\"data:image/png;base64, " + encodedData + "\">"
}
