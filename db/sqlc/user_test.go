package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:     fake.Food().Vegetable(),
		Role:     "client",
		Email:    fake.Internet().Email(),
		Password: "secret",
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.True(t, user.UpdatedAt.Valid)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	oldUser := createRandomUser(t)
	newUser, err := testStore.GetUser(context.Background(), oldUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, oldUser.Name, newUser.Name)
	require.Equal(t, oldUser.Email, newUser.Email)
	require.Equal(t, oldUser.Password, newUser.Password)
	require.WithinDuration(t, oldUser.UpdatedAt.Time, newUser.UpdatedAt.Time, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt.Time, newUser.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	oldUser := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), oldUser.ID)
	require.NoError(t, err)

	newUser, err := testStore.GetUser(context.Background(), oldUser.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, newUser)
}

func TestListUser(t *testing.T) {
	var lastUser User
	for i := 0; i < 10; i++ {
		lastUser = createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotEqual(t, lastUser.Name, user.Name)
	}
}

func TestUpdateUser(t *testing.T) {
	oldUser := createRandomUser(t)

	arg := UpdateUserParams{
		Name:     oldUser.Name,
		ID:       oldUser.ID,
		Password: oldUser.Password,
	}

	newUser, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, oldUser, newUser)
	require.WithinDuration(t, oldUser.CreatedAt.Time, newUser.CreatedAt.Time, time.Second)
}
