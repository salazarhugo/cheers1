package utils

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/api/option"
	"log"
)

type CustomContext struct {
	echo.Context
	neo4j.Driver
}

func GetDriver() neo4j.Driver {
	driver, err := neo4j.NewDriver(
		"neo4j+s://528a253a.databases.neo4j.io:7687",
		neo4j.BasicAuth("neo4j", "XRoQ6Lmz9QlFFTcwCWIWwR1o88hLfzV_HnP9mzDJuwc", ""))
	if err != nil {
		panic(err)
	}

	return driver
}

func GetSession(driver neo4j.Driver) neo4j.Session {
	return driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
}

func InitializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error utilsializing app: %v\n", err)
	}

	return app
}

func InitializeAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error utilsializing app: %v\n", err)
	}

	return client
}

func InitializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error utilsializing app: %v\n", err)
	}

	return app
}

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
