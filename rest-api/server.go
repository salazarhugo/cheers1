package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"rest-api/auth"
	"rest-api/event"
	"rest-api/post"
	"rest-api/story"
	"rest-api/user"
	"rest-api/utils"
)

func main() {
	e := echo.New()

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

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Event
	e.GET("/events", event.GetEvents)
	e.GET("/event/feed", event.GetEventFeed)
	e.POST("/event/create", event.CreateEvent)
	e.POST("/event/invite", event.InviteEvent)
	e.POST("/event/delete", event.DeleteEvent)
	e.POST("/event/interest", event.InterestEvent)
	e.POST("/event/uninterest", event.UninterestEvent)
	e.POST("/event/:eventId/going", event.GoingEvent)
	e.POST("/event/:eventId/ungoing", event.UngoingEvent)
	e.POST("/event/update", event.UpdateEvent)
	e.GET("/event/:eventId/going/list", event.GoingList)
	e.GET("/event/:eventId/interested/list", event.InterestedList)

	// Location
	e.POST("/updateLocation", user.UpdateLocation)
	e.GET("/locations", user.GetLocations)

	// Post
	e.GET("/posts", post.GetPosts)
	e.GET("/mapPosts", post.GetMapPosts)
	//e.GET("/posts/:postId", postDetails)
	e.GET("/posts/feed", post.PostFeed)
	e.POST("/posts/:postId/like", post.LikePost)
	e.POST("/posts/:postId/unlike", post.UnlikePost)
	e.POST("/posts/:postId/delete", post.DeletePost)
	e.POST("/posts/create", post.CreatePost)

	// Story
	e.GET("/mapStories", story.GetMapStories)
	e.GET("/stories/feed", story.StoryFeed)
	e.POST("/stories/create", story.CreateStory)
	e.POST("/stories/:storyId/delete", story.DeleteStory)
	e.POST("/stories/:storyId/seen", story.SeenStory)

	// User
	e.POST("/users/create", user.CreateUser)
	e.GET("/users/search/:query", user.SearchUsers)
	e.POST("/follow", user.FollowUser)
	e.POST("/unfollow", user.UnfollowUser)
	e.POST("/users/:userId/block", user.BlockUser)
	e.GET("/users/:userIdOrUsername", user.GetUser)
	e.GET("/users/available/:username", user.IsUsernameAvailable)
	e.POST("/users/tokens/:token", user.AddRegistrationToken)

	e.GET("/followers/list", user.FollowersList)
	e.GET("/following/list", user.FollowingList)

	e.GET("/users/activity", user.GetActivity)

	// Auth
	e.GET("/token", auth.CreateCustomToken)

	e.Logger.Fatal(e.Start(":" + port))
}
