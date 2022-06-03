package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"os"
)

type Properties struct {
	PhotoUrl string `json:"PhotoUrl"`
	Username string `json:"username"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type CustomContext struct {
	echo.Context
	neo4j.Driver
}

func getDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver(
		"neo4j+s://528a253a.databases.neo4j.io:7687",
		neo4j.BasicAuth("neo4j", "XRoQ6Lmz9QlFFTcwCWIWwR1o88hLfzV_HnP9mzDJuwc", ""))
	if err != nil {
		panic(err)
	}

	return driver
}

func getSession(driver neo4j.Driver) neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, getDriver()}
			return next(cc)
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://web-cheers.web.app", "http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.Use(AuthMiddleware)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Event
	e.GET("/events", getEvents)
	e.GET("/event/feed", getEventFeed)
	e.POST("/event/create", createEvent)
	e.POST("/event/invite", inviteEvent)
	e.POST("/event/delete", deleteEvent)
	e.POST("/event/interest", interestEvent)
	e.POST("/event/uninterest", uninterestEvent)
	e.POST("/event/:eventId/going", goingEvent)
	e.POST("/event/:eventId/ungoing", ungoingEvent)
	e.POST("/event/update", updateEvent)

	// Location
	e.POST("/updateLocation", updateLocation)
	e.GET("/locations", getLocations)

	// Post
	e.GET("/posts", getPosts)
	e.GET("/mapPosts", getMapPosts)
	//e.GET("/posts/:postId", postDetails)
	e.GET("/posts/feed", postFeed)
	e.POST("/posts/:postId/like", likePost)
	e.POST("/posts/:postId/unlike", unlikePost)
	e.POST("/posts/:postId/delete", deletePost)
	e.POST("/posts/create", createPost)

	// Story
	e.GET("/mapStories", getMapStories)
	e.GET("/stories/feed", storyFeed)
	e.POST("/stories/create", createStory)
	e.POST("/stories/:storyId/delete", deleteStory)
	e.POST("/stories/:storyId/seen", seenStory)

	// User
	e.POST("/users/create", createUser)
	e.GET("/users/search/:query", searchUsers)
	e.POST("/follow", followUser)
	e.POST("/unfollow", unfollowUser)
	e.POST("/users/:userId/block", blockUser)
	e.GET("/users/:userIdOrUsername", getUser)
	e.GET("/users/available/:username", isUsernameAvailable)
	e.POST("/users/tokens/:token", addRegistrationToken)

	e.GET("/followers/list", followersList)
	e.GET("/following/list", followingList)

	e.GET("/users/activity", getActivity)

	// Auth
	e.GET("/token", createCustomToken)

	e.Logger.Fatal(e.Start(":" + port))
}
