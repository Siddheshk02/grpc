package server

import (
	"context"
	"testing"

	"github.com/Siddheshk02/grpc/models"
	"github.com/Siddheshk02/grpc/pb"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	srv := NewUserServiceServer()
	srv.users[1] = models.User{
		ID:      1,
		FName:   "Steve",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.8,
		Married: true,
	}

	res, err := srv.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Steve", res.User.Fname)
}

func TestListUsers(t *testing.T) {
	srv := NewUserServiceServer()
	srv.users[1] = models.User{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	srv.users[2] = models.User{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false}

	res, err := srv.ListUsers(context.Background(), &pb.ListUsersRequest{Ids: []int32{1, 2}})
	assert.NoError(t, err)
	assert.Len(t, res.Users, 2)
}

func TestSearchUsers(t *testing.T) {
	srv := NewUserServiceServer()
	srv.users[1] = models.User{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	srv.users[2] = models.User{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false}

	res, err := srv.SearchUsers(context.Background(), &pb.SearchUsersRequest{City: "LA", Phone: 1234567890, Married: true})
	assert.NoError(t, err)
	assert.Len(t, res.Users, 2)
}
