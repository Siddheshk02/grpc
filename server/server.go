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
	users map[int32]models.User
	mu    sync.Mutex
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: make(map[int32]models.User),
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.Id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{User: &pb.User{
		Id:      user.ID,
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}}, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User
	for _, id := range req.Ids {
		if user, exists := s.users[id]; exists {
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

	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
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
