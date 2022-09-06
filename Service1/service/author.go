package service

import (
	"context"
	"service1/proto"
	"service1/repository"
)

func (s *Server) FindBookByIdAuthor(ctx context.Context, request *proto.FindBookByIdAuthorRequest) (*proto.BooksResponse, error) {
	response, err := repository.FindBookByIdAuthor(request)
	res := &proto.BooksResponse{
		BooksResponse: response,
	}
	return res, err
}
