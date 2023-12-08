package user_service

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up your environment variables for testing
	dsn := "projects/cheers-a275e/instances/free-cheers/databases/party"
	os.Setenv("NEO4J_URI", "neo4j+s://86c8796d.databases.neo4j.io")
	os.Setenv("NEO4J_PASSWORD", "z1nNlYESkGiq-MelKNPP3HBYeh23jWZLA6faaVA1RpM")
	os.Setenv("SPANNER_DSN", dsn)

	// Run the tests
	result := m.Run()

	// Clean up or perform any additional actions after tests if needed

	// Exit with the test result
	os.Exit(result)
}
