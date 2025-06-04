package handler

import (
	"context"
	"math/rand"
	"time"
	"user-server/basic/inits"
	__ "user-server/proto"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Sendsms(_ context.Context, in *__.SendsmsRequest) (*__.SendsmsResponse, error) {
	code := rand.Intn(9000) + 1000
	inits.RedisClient.Set(context.Background(), "sendsms"+in.Tel+in.Score, code, time.Minute*5)
	return &__.SendsmsResponse{}, nil
}
