package service

import (
	"fmt"
	"net"
	"service1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func StartService(listener net.Listener) {
	srv := grpc.NewServer()
	proto.RegisterAddAuthorServiceServer(srv, &Server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		fmt.Println(e)
	}
}
