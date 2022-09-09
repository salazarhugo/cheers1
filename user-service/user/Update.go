package user

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"salazar/cheers/user/userpb"
	"salazar/cheers/user/utils"
)

// UpdateUser PATCH /users
func UpdateUser(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	userId := cc.Get("userId").(string)

	user := userpb.UpdateUserRequest{}

	err := json.NewDecoder(cc.Request().Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding request body.")
		return cc.NoContent(http.StatusBadRequest)
	}

	cypher := ` MATCH (u: User {id: $userId})
                SET u += $user`

	params := map[string]interface{}{
		"userId": userId,
		"user":   structs.Map(user),
	}

	_, err = session.Run(cypher, params)

	if err != nil {
		log.Println(err)
		return err
	}

	return cc.NoContent(http.StatusOK)
}
