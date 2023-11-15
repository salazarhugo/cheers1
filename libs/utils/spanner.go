package utils

import (
	spannergorm "github.com/googleapis/go-gorm-spanner"
	_ "github.com/googleapis/go-sql-spanner"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"os"
)

func GetSpanner() *gorm.DB {
	db, err := gorm.Open(
		spannergorm.New(
			spannergorm.Config{
				DriverName: "spanner",
				DSN:        os.Getenv("SPANNER_DSN"),
			},
		),
		&gorm.Config{PrepareStmt: true},
	)
	if err != nil {
		log.Fatalf("failed to open database connection: %v\n", err)
	}

	return db
}
