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
			Id:    "_KVP1NWX4GEMhOakUGwbdg",
			RawId: "X0tWUDFOV1g0R0VNaE9ha1VHd2JkZw",
			Type:  "public-key",
			Response: &auth.Response{
				AttestationObject: "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YViUD59Ji51NmDKORzsHeTZd1BCLAEcHsprui1AidZwIGtxdAAAAAOqbjWZNAR0hPOS2tIy1ddQAEPylT9TVl-BhDITmpFBsG3alAQIDJiABIVggL0iMlHpeG9yywiuU65N1UFBhgAvHAIQoYw33mgBqDNEiWCDXwzUsffuCN9Z9wxgYvRNHrTXLtDenlW4OZZGxlvuYeA",
				ClientData_JSON:   "eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoiSHIzZDdRSkphdDJzVHd4NjN1NlhpMHZZdUlKOW9DYkNjdVc0VDZmamV6YyIsIm9yaWdpbiI6ImFuZHJvaWQ6YXBrLWtleS1oYXNoOkJrY2o4ZXQyMVcyamk3SC10dWdoQmkxNVdDMVhrNU5NTlY5cnpHZmQ0b0kiLCJhbmRyb2lkUGFja2FnZU5hbWUiOiJjb20uc2FsYXphci5jaGVlcnMifQ",
			},
		},
		UserId:    13729502662429043192,
		Challenge: "Hr3d7QJJat2sTwx63u6Xi0vYuIJ9oCbCcuW4T6fjezc",
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
