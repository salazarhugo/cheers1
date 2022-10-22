package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	storypb "github.com/salazarhugo/cheers1/genproto/cheers/type/story"
)

type StoryRepository interface {
	CreateStory(userID string, story *storypb.Story) (string, error)
	GetStory(userID string, storyID string) (*pb.StoryResponse, error)
	UpdateStory(story *storypb.Story) (*pb.StoryResponse, error)
	DeleteStory(id string) error

	FeedStory(userID string, request *pb.FeedStoryRequest) (*pb.FeedStoryResponse, error)
	ViewStory(userID string, storyID string) (*pb.ViewStoryResponse, error)
	LikeStory(userID string, storyID string) (*pb.LikeStoryResponse, error)
	UnlikeStory(userID string, storyID string) (*pb.UnlikeStoryResponse, error)
}

type storyRepository struct {
	driver neo4j.Driver
}

func NewStoryRepository(driver neo4j.Driver) StoryRepository {
	return &storyRepository{driver: driver}
}
