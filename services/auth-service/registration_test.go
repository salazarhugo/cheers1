package auth_service

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/app"
	"testing"
)

func TestRegistration(t *testing.T) {

	request := &auth.FinishRegistrationRequest{
		Email:    "hugobrock74@gmail.com",
		Username: "lars",
		Passkey: &auth.CreatePasskeyResponseData{
			Id:    "GpTQg7zMb4w0iPSqt62NuQ",
			RawId: "R3BUUWc3ek1iNHcwaVBTcXQ2Mk51UQ",
			Type:  "public-key",
			Response: &auth.Response{
				AttestationObject: "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViUD59Ji51NmDKORzsHeTZd1BCLAEcHsprui1AidZwIGtxdAAAAAOqbjWZNAR0hPOS2tIy1ddQAEBqU0IO8zG-MNIj0qretjbmlAQIDJiABIVggRPfskV5gFmfzPgSqVwkQoxcfb_OsTGiPRYHDWR2il8wiWCByRYED0i3XrN44mtAspmK51Eoth-v3DulNwA69VHsfZA",
				ClientData_JSON:   "eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoia1hCM0N1cTFOMEw1bUZjYkVkWXVSeGs2Vlp4X3VhTEpVS3JQVmZ1Wm81WSIsIm9yaWdpbiI6ImFuZHJvaWQ6YXBrLWtleS1oYXNoOkJrY2o4ZXQyMVcyamk3SC10dWdoQmkxNVdDMVhrNU5NTlY5cnpHZmQ0b0kiLCJhbmRyb2lkUGFja2FnZU5hbWUiOiJjb20uc2FsYXphci5jaGVlcnMifQ",
			},
		},
	}

	server := app.NewServer()
	_, err := server.FinishRegistration(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}
}
