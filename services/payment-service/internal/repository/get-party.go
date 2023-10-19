package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/party"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func GetParty(partyId string) (*partypb.Party, error) {
	ctx := context.Background()

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}
	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	conn, err := grpc.DialContext(ctx, "party-service-r3a2dr4u4a-nw.a.run.app:443",
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithUnaryInterceptor(utils.CloudRunInterceptor),
	)
	defer conn.Close()

	client := party.NewPartyServiceClient(conn)

	party, err := client.GetParty(ctx, &party.GetPartyRequest{PartyId: partyId})
	if err != nil {
		return nil, err
	}

	return party.GetParty(), nil
}
