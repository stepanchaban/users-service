package grpc

import (
	"context"

	userpb "github.com/stepanchaban/project-prot/proto/user"
	"github.com/stepanchaban/users-service/internal/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.CreateUser(user.UserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       u.ID,
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	u, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:       u.ID,
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

func (h *Handler) ListUsers(ctx context.Context, _ *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:       u.ID,
			Email:    u.Email,
			Password: u.Password,
		})
	}

	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u, err := h.svc.UpdateUser(req.Id, user.UserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       u.ID,
			Email:    u.Email,
			Password: u.Password,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	err := h.svc.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}