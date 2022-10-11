package service

import (
	"context"
	"service1/proto"
	"service1/repository"
)

type Server struct {
	repository repository.AuthService1Repo
}

func NewService(repository repository.AuthService1Repo) *Server {

	return &Server{
		repository: repository,
	}
}

func (s *Server) FindBookByIdAuthor(ctx context.Context, request *proto.FindBookByIdAuthorRequest) (*proto.BooksResponse, error) {
	response, err := s.repository.FindBookByIdAuthor(request)

	res := &proto.BooksResponse{
		BooksResponse: response,
	}
	return res, err
}

func (s *Server) LoginGPRC(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	response, err := s.repository.LoginGRPC(request)
	return response, err
}

func (s *Server) Logout(ctx context.Context, request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	res, err := s.repository.Logout(request)
	if err != nil {
		return &proto.LogoutResponse{Status: "false"}, err
	}
	return res, nil
}
