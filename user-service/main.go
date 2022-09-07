package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"salazar/cheers/user/auth"
	"salazar/cheers/user/user"
	"salazar/cheers/user/utils"
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

	e.Use(auth.UserInfoMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	e.GET("/", user.GetUser)
	e.PATCH("/", user.UpdateUser)
	//e.POST("/", user.CreateUser)
	//e.GET("/tickets", user.GetUserTickets)
	//e.GET("/search/:query", user.SearchUsers)
	//e.POST("/follow", user.FollowUser)
	//e.POST("/unfollow", user.UnfollowUser)
	//e.POST("/:userId/block", user.BlockUser)
	//e.GET("/available/:username", user.IsUsernameAvailable)
	//e.GET("/suggestions", user.UserSuggestions)
	//e.POST("/tokens/:token", user.AddRegistrationToken)
	e.Logger.Fatal(e.Start(":" + port))
}
