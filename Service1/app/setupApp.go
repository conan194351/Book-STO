package app

import (
	"fmt"
	"net"
	"service1/config"
	"service1/redis"
	"service1/service"
)

func RunApp() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
	}
	config.InitDatabase()
	redis.NewResdisClient()
	service.StartService(listener)
}
