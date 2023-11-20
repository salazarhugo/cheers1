package auth_service

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/app"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/repository"
	"testing"
)

func TestRegistration(t *testing.T) {
	//Create a mock repository
	repo := repository.NewRepository()

	server := app.NewServer()
	request := &auth.FinishRegistrationRequest{
		Email:    "hugobrock74@gmail.com",
		Username: "lars",
		Passkey: &auth.CreatePasskeyResponseData{
			Id:    "6bmjCKxTBP2xazq0Tn873g",
			RawId: "NmJtakNLeFRCUDJ4YXpxMFRuODczZw==",
			Type:  "public-key",
			Response: &auth.Response{
				AttestationObject: "bzJObWJYUmtibTl1WldkaGRIUlRkRzEwb0doaGRYUm9SR0YwWVZpVUQ1OUppNTFObURLT1J6c0hlVFpkMUJDTEFFY0hzcHJ1aTFBaWRad0lHdHhkQUFBQUFPcWJqV1pOQVIwaFBPUzJ0SXkxZGRRQUVPbTVvd2lzVXdUOXNXczZ0RTVfTzk2bEFRSURKaUFCSVZnZzRUSE95OWxqZnhCT2JpNXI2azl2QXpKSUpBdmNta0gycHR3eWlvSV9ZdWtpV0NEajZoOW5Dd1l5dkx0R1RiM2pmMEJBU3pGSXR0VDdjYUVJa2VTQkRtbTZZQQ==bzJObWJYUmtibTl1WldkaGRIUlRkRzEwb0doaGRYUm9SR0YwWVZpVUQ1OUppNTFObURLT1J6c0hlVFpkMUJDTEFFY0hzcHJ1aTFBaWRad0lHdHhkQUFBQUFPcWJqV1pOQVIwaFBPUzJ0SXkxZGRRQUVPbTVvd2lzVXdUOXNXczZ0RTVfTzk2bEFRSURKaUFCSVZnZzRUSE95OWxqZnhCT2JpNXI2azl2QXpKSUpBdmNta0gycHR3eWlvSV9ZdWtpV0NEajZoOW5Dd1l5dkx0R1RiM2pmMEJBU3pGSXR0VDdjYUVJa2VTQkRtbTZZQQ==",
				ClientDataJson:    "",
			},
		},
	}
	server.FinishRegistration(context.Background(), request)
	_, err := repo.CreateCredential(
		"user1",
		"publickey23r52582",
	)
	if err != nil {
		t.Error("failed to create credential: ", err)
		return
	}
}
