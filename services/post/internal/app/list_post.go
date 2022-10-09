package app

//	v1 "github.com/salazarhugo/cheers1/genproto/cheers/api/v1"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//)
//
//func (m *MicroserviceServer) ListParty(
//	ctx context.Context,
//	request *v1.CreatePartyRequest,
//) (*party.Party, error) {
//	userID, err := GetUserId(ctx)
//	if err != nil {
//		return nil, status.Error(codes.Internal, "Failed retrieving userID")
//	}
//
//	partyReq := request.GetParty()
//	if partyReq == nil {
//		return nil, status.Error(codes.InvalidArgument, "party parameter can't be nil")
//	}
//
//	partyReq.HostId = userID
//
//	err = m.partyRepository.CreateParty(*partyReq)
//	if err != nil {
//		return nil, status.Error(codes.Internal, "failed to create party")
//	}
//
//	return partyReq, nil
//}
