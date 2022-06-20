package comment

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"rest-api/usereventpb"
	"rest-api/utils"
	"time"
)

type Comment struct {
	Id     string `json:"id" structs:"id"`
	Text   string `json:"text" structs:"text"`
	PostId string `json:"postId" structs:"postId"`
}

func CreateComment(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)
	comment := Comment{}

	err := json.NewDecoder(c.Request().Body).Decode(&comment)
	if err != nil {
		return err
	}

	state := &usereventpb.UserEvent{
		Type:   usereventpb.UserEventType_COMMENT,
		UserId: userId,
		PostId: comment.PostId,
		Time:   time.Now().Unix(),
	}
	err = utils.PublishProtoMessages(state)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
