package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"rest-api/auth"
	"rest-api/comment"
	"rest-api/event"
	"rest-api/post"
	"rest-api/story"
	"rest-api/user"
	"rest-api/utils"
)

func main() {
	e := echo.New()

	log.SetFlags(0)
	//client, err := logging.NewClient(context.Background(), projectID)
	//if err != nil {
	//	log.Fatalf("Failed to create client: %v", err)
	//}

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

	// Party
	e.GET("/events", event.GetEvents)
	e.GET("/party/feed", event.GetPartyFeed)
	e.GET("/party/my", event.GetMyParties)
	e.POST("/party/create", event.CreateParty)
	e.POST("/event/invite", event.InviteEvent)
	e.DELETE("/party/:partyId", event.DeleteParty)
	e.GET("/party/:partyId", event.GetParty)
	e.POST("/party/interest", event.InterestParty)
	e.POST("/party/uninterest", event.UninterestParty)
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
	e.DELETE("/posts/:id", post.DeletePost)
	e.GET("/posts/:userId", post.GetUserPosts)
	//e.GET("/posts/history", user.GetHistoryPosts)
	e.GET("/mapPosts", post.GetMapPosts)
	e.GET("/posts/feed", post.PostFeed)
	e.GET("/posts/:postId/members", post.PostMembers)
	e.POST("/posts/:postId/like", post.LikePost)
	e.POST("/posts/:postId/unlike", post.UnlikePost)
	e.POST("/posts/create", post.CreatePost)

	// Comments
	e.POST("/comment", comment.CreateComment)

	// Story
	e.GET("/stories/:username", story.GetUserStory)
	e.GET("/mapStories", story.GetMapStories)
	e.GET("/stories/feed", story.StoryFeed)
	e.POST("/stories/create", story.CreateStory)
	e.POST("/stories/:storyId", story.DeleteStory)
	e.POST("/stories/:storyId/seen", story.SeenStory)

	// User
	e.PATCH("/users", user.UpdateUser)
	e.POST("/users/create", user.CreateUser)
	e.GET("/users/tickets", user.GetUserTickets)
	e.GET("/users/search/:query", user.SearchUsers)
	e.POST("/follow", user.FollowUser)
	e.POST("/unfollow", user.UnfollowUser)
	e.POST("/users/:userId/block", user.BlockUser)
	e.GET("/users/:userIdOrUsername", user.GetUser)
	e.GET("/users/available/:username", user.IsUsernameAvailable)
	e.GET("/users/suggestions", user.UserSuggestions)
	e.POST("/users/tokens/:token", user.AddRegistrationToken)

	e.GET("/followers/list", user.FollowersList)
	e.GET("/following/list", user.FollowingList)

	e.GET("/users/activity", user.GetActivity)

	// Auth
	e.GET("/token", auth.CreateCustomToken)

	e.Logger.Fatal(e.Start(":" + port))
}
