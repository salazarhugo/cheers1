package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/story-service/internal/repository"
	"sync"
)

type Server struct {
	pb.UnimplementedStoryServiceServer
	mu              sync.Mutex
	storyRepository repository.StoryRepository
}

func (s *Server) SaveStory(ctx context.Context, request *pb.SaveStoryRequest) (*pb.SaveStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UnsaveStory(ctx context.Context, request *pb.UnsaveStoryRequest) (*pb.UnsaveStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer() *Server {
	storyRepository := repository.NewStoryRepository(utils.GetDriver())

	return &Server{
		UnimplementedStoryServiceServer: pb.UnimplementedStoryServiceServer{},
		mu:                              sync.Mutex{},
		storyRepository:                 storyRepository,
	}
}
