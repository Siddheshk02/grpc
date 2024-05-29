package server

import (
	"context"
	"errors"
	"sync"

	"github.com/Siddheshk02/grpc/models"
	"github.com/Siddheshk02/grpc/pb"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	users []models.User
	mu    sync.Mutex
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: []models.User{
			{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
			{ID: 3, FName: "Sid", City: "MU", Phone: 9567843210, Height: 5.7, Married: true},
			{ID: 4, FName: "Satoshi", City: "JP", Phone: 6789543210, Height: 5.5, Married: false},
			{ID: 5, FName: "Vitalik", City: "WR", Phone: 9876012345, Height: 5.4, Married: true},
			{ID: 6, FName: "Elon", City: "CA", Phone: 9823456710, Height: 6.0, Married: false},
			{ID: 7, FName: "Mark", City: "CA", Phone: 2345678910, Height: 6.1, Married: true},
		},
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.ID == req.Id {
			return &pb.GetUserResponse{User: &pb.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}}, nil
		}
	}
	return nil, errors.New("user not found")

}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.ID == id {
				users = append(users, &pb.User{
					Id:      user.ID,
					Fname:   user.FName,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
				break
			}
		}
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User
	for _, user := range s.users {
		if (req.Fname == "" || user.FName == req.Fname) &&
			(req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			users = append(users, &pb.User{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}

	return &pb.SearchUsersResponse{Users: users}, nil
}
