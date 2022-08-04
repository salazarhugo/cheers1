package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"salazar/cheers/payment/auth"
	"salazar/cheers/payment/endpoints"
	"salazar/cheers/payment/utils"
)

func main() {
	e := echo.New()

	log.SetFlags(0)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &utils.CustomContext{Context: c, Driver: utils.GetDriver()}
			return next(cc)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://web-cheers.web.app", "http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(auth.AuthMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	e.POST("/create-payment-intent", endpoints.CreatePaymentIntent)

	e.Logger.Fatal(e.Start(":" + port))
}
