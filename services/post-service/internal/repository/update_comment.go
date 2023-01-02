package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/type/user"
	"github.com/salazarhugo/cheers1/libs/utils"
)

func (p *postRepository) UpdatePostLastComment(
	user *user.UserItem,
	comment *comment.Comment,
) error {
	session := p.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher, err := utils.GetCypher("internal/queries/UpdatePostComment.cql")
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"postID":                comment.PostId,
		"lastCommentText":       comment.Text,
		"lastCommentUsername":   user.Username,
		"lastCommentCreateTime": comment.CreateTime,
	}

	_, err = session.Run(*cypher, params)
	if err != nil {
		return err
	}

	// TODO Publish Update Post Event

	return nil
}
