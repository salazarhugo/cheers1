package comment

import (
	"context"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"net/http"
	"rest-api/user"
	"rest-api/usereventpb"
	"rest-api/utils"
	"time"
)

type Comment struct {
	Id       string `json:"id" structs:"id"`
	AuthorId string `json:"authorId" structs:"authorId"`
	Text     string `json:"text" structs:"text"`
	PostId   string `json:"postId" structs:"postId"`
	Username string `json:"username" structs:"username"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Verified bool   `json:"verified" structs:"verified"`
	Created  int64  `json:"created" structs:"created"`
}

func CreateComment(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)

	ctx := context.Background()
	app := utils.InitializeAppDefault()
	client, err := app.Firestore(ctx)

	comment := Comment{}
	err = json.NewDecoder(c.Request().Body).Decode(&comment)
	if err != nil {
		return err
	}
	userCard := user.GetUserCard(c, userId)

	commentDoc := client.Collection("comments").NewDoc()

	comment.Id = commentDoc.ID
	comment.Created = time.Now().Unix()
	comment.AuthorId = userId
	comment.Username = userCard.Username
	comment.Verified = userCard.Verified
	comment.Avatar = userCard.Avatar

	commentDoc.Set(ctx, structs.Map(comment))

	state := &usereventpb.UserEvent{
		Type:   usereventpb.UserEventType_COMMENT,
		UserId: userId,
		PostId: comment.PostId,
		Time:   comment.Created,
	}
	err = utils.PublishProtoMessages(state)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
