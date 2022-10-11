package service

import (
	"fmt"
	"net"
	"service1/config"
	"service1/proto"
	"service1/redis"
	"service1/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartService(listener net.Listener) {
	srv := grpc.NewServer()
	var server = NewService(repository.NewAuthService1Repo(config.DB, redis.RDB))
	proto.RegisterAddAuthorServiceServer(srv, server)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		fmt.Println(e)
	}
}
