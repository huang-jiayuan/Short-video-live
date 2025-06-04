package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-server/basic/config"
	"user-server/basic/inits"
	"user-server/handler"
	__ "user-server/proto"
)

func main() {
	config.Inits()
	inits.InitMysql()
	inits.RedisInit()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &handler.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
