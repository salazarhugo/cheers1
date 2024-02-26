package tests

import (
	"github.com/salazarhugo/cheers1/libs/utils/models"
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"testing"
)

func CreateRandomUser(t *testing.T) *models.User {
	return tests.CreateRandomUser(t)
}
