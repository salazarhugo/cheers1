package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"salazar/cheers/payment/utils"
)

type Ticket struct {
	ID string `query:"id"`
}

func ValidateTicket(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()

	// in the handler for ticket/validate?id=<ticketId>
	var ticketReq Ticket
	err := c.Bind(&ticketReq)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request: Missing ticket id")
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cheers-a275e")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	doc, err := client.Collection("tickets").Doc(ticketReq.ID).Get(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Ticket not found")
	}
	ticket := doc.Data()

	return cc.JSON(http.StatusOK, map[string]interface{}{
		"valid":       true,
		"name":        ticket["name"],
		"description": ticket["description"],
		"message":     ticket["name"],
	})
}
