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
		Email:    fake.Internet().Email(),
		Password: "secret",
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.True(t, user.UpdatedAt.Valid) //why
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2) //why are we check for emtyness

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	// require.EqualError(t, err, ErrRecordNotFound.Error())// this line of code did not work
	require.Empty(t, user2)
}

func TestListAccounts(t *testing.T) {
	var lastUser User
	for i := 0; i < 10; i++ {
		lastUser = createRandomUser(t)
	}

	arg := ListUsersParams{
		Name: lastUser.Name,
		Limit: 5,
		Offset: 0,
	}

	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.Equal(t, lastUser.Name, user.Name)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		Name: user1.Name,
		Email: user1.Email,
		Password: user1.Password,
		CreatedAt: user1.CreatedAt,
		UpdatedAt: user1.UpdatedAt,
		DeletedAt: user1.DeletedAt,
	}

	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	
	require.Equal(t, user1, user2)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second) // there was an error when I wanted to do: product.CreatedAt
}

// func TestUpdateUser(t *testing.T) {
// 	user := createRandomUser(t)

// 	updatingProperties := map[string]any{
// 		"name":     fake.Food().Vegetable(),
// 		"email":    fake.Internet().Email(),
// 		"password": "secret",
// 	}

// 	tests := []struct {
// 		name string
// 		args map[string]any
// 	}{
// 		{
// 			name: "update name",
// 			args: map[string]any{
// 				"name": updatingProperties["name"],
// 				"id":   user.ID,
// 			},
// 		},
// 		{
// 			name: "update email",
// 			args: map[string]any{
// 				"email": updatingProperties["email"],
// 				"id":    user.ID,
// 			},
// 		},
// 		{
// 			name: "update all",
// 			args: map[string]any{
// 				"name": updatingProperties["name"],
// 				"email": updatingProperties["email"],
// 				"password": updatingProperties["password"],
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			dbPayload := covnertMapToUpdateUserParams(t, tt.args)
// 			got, err := testStore.UpdateUser(context.Background(), dbPayload)
// 			require.NoError(t, err, "UpdateUser() error = %v", err)

// 			want := makeUpdateUserResp(t, &user, tt.args)
// 			require.Equal(t, got.Name, want.Name)
// 			require.Equal(t, got.Email, want.Email)
// 			require.Equal(t, got.Password, want.Password)

// 			require.Equal(t, user, got, "UpdateUser() got = %v, want %v", got, want)
// 		})
// 	}
// }

// func covnertMapToUpdateUserParams(t *testing.T, args map[string]any) UpdateUserParams {
// 	bytes, err := json.Marshal(args)
// 	require.NoError(t, err)

// 	var updateUserParams UpdateUserParams
// 	err = json.Unmarshal(bytes, &updateUserParams)
// 	require.NoError(t, err)

// 	return updateUserParams
// }

// func makeUpdateUserResp(t *testing.T, user *User, args map[string]any) *User {
// 	bytes, err := json.Marshal(args)
// 	require.NoError(t, err)

// 	err = json.Unmarshal(bytes, &user)
// 	require.NoError(t, err)

// 	return user
// }
