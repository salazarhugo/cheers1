package endpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"net/http"
	"os"
	"salazar/cheers/payment/utils"
)

func ValidateTicket(c echo.Context) error {
	cc := c.(*utils.CustomContext)
	session := utils.GetSession(cc.Driver)
	defer session.Close()
	stripe.Key = os.Getenv("STRIPE_SK")

	return cc.NoContent(http.StatusOK)
}
