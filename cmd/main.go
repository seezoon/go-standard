package main

import (
	log "github.com/sirupsen/logrus"
	"go-standard/entity/base"
	"go-standard/pb"
	"go-standard/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service.NewUserService())
	lis, err := net.Listen(base.ConfInstance.Grpc.Network, base.ConfInstance.Grpc.Listen)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
