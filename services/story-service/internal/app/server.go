package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/story-service/internal/repository"
	"sync"
)

type Server struct {
	pb.UnimplementedStoryServiceServer
	mu              sync.Mutex
	storyRepository repository.StoryRepository
}

func (s *Server) FeedStory(ctx context.Context, request *pb.FeedStoryRequest) (*pb.FeedStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) LikeStory(ctx context.Context, request *pb.LikeStoryRequest) (*pb.LikeStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UnlikeStory(ctx context.Context, request *pb.UnlikeStoryRequest) (*pb.UnlikeStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SaveStory(ctx context.Context, request *pb.SaveStoryRequest) (*pb.SaveStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UnsaveStory(ctx context.Context, request *pb.UnsaveStoryRequest) (*pb.UnsaveStoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedStoryServiceServer() {
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
