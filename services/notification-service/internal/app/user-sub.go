package app

//func UserEventPubSub(c echo.Context) error {
//	var m PubSubMessage
//	body, err := io.ReadAll(c.Request().Body)
//	if err != nil {
//		log.Printf("ioutil.ReadAll: %v", err)
//		return c.String(http.StatusBadRequest, "Bad Request")
//	}
//	if err := json.Unmarshal(body, &m); err != nil {
//		log.Printf("json.Unmarshal: %v", err)
//		return c.String(http.StatusBadRequest, "Bad Request")
//	}
//
//	userevent := &usereventpb.UserEvent{}
//	err = proto.Unmarshal(m.Message.Data, userevent)
//	if err != nil {
//		log.Fatal(err)
//		return err
//	}
//
//	switch userevent.Type {
//	case usereventpb.UserEventType_POST_LIKE:
//		likePostNotification(c, userevent)
//	case usereventpb.UserEventType_CREATE_POST:
//		postNotification(c, userevent)
//	case usereventpb.UserEventType_FOLLOW:
//		followNotification(c, userevent)
//	case usereventpb.UserEventType_COMMENT:
//		commentNotification(c, userevent)
//	default:
//	}
//
//	return c.NoContent(http.StatusOK)
//}
