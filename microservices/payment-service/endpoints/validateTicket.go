package endpoints

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"os"
	"salazar/cheers/payment/proto/userpb"
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

	ticketDoc, err := client.Collection("tickets").Doc(ticketReq.ID).Get(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Ticket not found")
	}

	ticketMap := ticketDoc.Data()
	ticket, err := utils.MapToTicket(ticketMap)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to Unmarshal ticket")
	}

	conn, err := grpc.Dial(os.Getenv("GATEWAY_DOMAIN")+":443", grpc.WithInsecure())
	log.Println(conn.Target())
	if err != nil {
		log.Println("DIal")
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+os.Getenv("TOKEN"))
	userClient := userpb.NewUserServiceClient(conn)
	resp, err := userClient.GetUser(ctx, &userpb.GetUserRequest{})
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := resp.GetUser()

	// Check if ticket was already validated
	if ticketMap["validated"] == true {
		return cc.JSON(http.StatusOK, map[string]interface{}{
			"valid":            false,
			"name":             ticket.Name,
			"description":      ticket.Description,
			"message":          "Ticket already validated",
			"price":            ticket.Price,
			"reservation_name": user.Name,
			"party_name":       "Summer Breeze",
			"party_banner":     "Summer Breeze",
		})
	}

	batch := client.Batch()

	batch.Update(ticketDoc.Ref, []firestore.Update{
		{Path: "validated", Value: true},
	})

	_, err = batch.Commit(ctx)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = utils.PublishPayment(user.Email)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return cc.JSON(http.StatusOK, map[string]interface{}{
		"valid":           true,
		"name":            ticketMap["name"],
		"description":     ticketMap["description"],
		"message":         ticketMap["name"],
		"price":           ticketMap["price"],
		"reservationName": user.Name,
		"provider":        "Summer Breeze",
	})
}
