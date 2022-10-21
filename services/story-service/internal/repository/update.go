package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/story/v1"
	storypb "github.com/salazarhugo/cheers1/genproto/cheers/type/story"
)

func (p *storyRepository) UpdateStory(
	story *storypb.Story,
) (*pb.StoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
