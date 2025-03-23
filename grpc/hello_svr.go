package main

import (
	"context"
	"fmt"
	proto "github.com/xiuluoyidao/demo/proto/github.com/xiuluoyidao/demo/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &HelloServiceServer{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// HelloServiceServer 是服务实现
type HelloServiceServer struct{}

func (s *HelloServiceServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: fmt.Sprintf("我是航仔！Hello, %s!", req.Name),
	}, nil
}

//func (s *HelloServiceServer) mustEmbedUnimplementedGreeterServer() {
//	return
//}
