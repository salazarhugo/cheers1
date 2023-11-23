package app

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"log"
)

func NewWebAuthn() (*webauthn.WebAuthn, error) {
	allowedOrigins := []string{
		"https://cheers.social",
		"android:apk-key-hash:Bkcj8et21W2ji7H-tughBi15WC1Xk5NMNV9rzGfd4oI",
	}

	wconfig := &webauthn.Config{
		RPDisplayName: "Cheers",        // Display Name for your site
		RPID:          "cheers.social", // Generally the FQDN for your site
		RPOrigins:     allowedOrigins,  // The origin URLs allowed for WebAuthn requests
	}

	webAuthn, err := webauthn.New(wconfig)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return webAuthn, nil
}
