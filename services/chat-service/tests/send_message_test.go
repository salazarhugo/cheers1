package tests

import (
	"github.com/google/uuid"
	"github.com/salazarhugo/cheers1/libs/utils/tests"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendMessage(t *testing.T) {
	user1 := tests.CreateRandomUser(t)
	user2 := tests.CreateRandomUser(t)
	repo := repository.NewChatRepository(nil)

	room, err := repo.CreateRoom("yo", []string{user1.UserId, user2.UserId})
	require.NoError(t, err)
	t.Log(room)

	_, err = repo.SendMessage(
		uuid.NewString(),
		user1.UserId,
		room.Id,
		"Hello there",
		"",
	)
	require.NoError(t, err)

	err = repo.DeleteRooms(user1.UserId)
	require.NoError(t, err)
}
