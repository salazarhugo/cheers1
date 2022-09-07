package user

//// POST("/users/create")
//func CreateUser(c echo.Context) error {
//	cc := c.(*utils.CustomContext)
//	session := utils.GetSession(cc.Driver)
//	defer session.Close()
//
//	userReq := userpb.User{}
//	userId := cc.Get("userId").(string)
//	err := json.NewDecoder(cc.Request().Body).Decode(&userReq)
//
//	if err != nil {
//		return err
//	}
//
//	userReq.Id = userId
//	if _, err := session.WriteTransaction(addUserNodeTxFunc(userReq)); err != nil {
//		log.Print(err)
//		return err
//	}
//
//	var user interface{}
//	if user, err = session.ReadTransaction(matchUserNodeTxFunc(userReq.Id)); err != nil {
//		log.Print(err)
//		return err
//	}
//
//	return cc.JSON(http.StatusOK, user)
//}
