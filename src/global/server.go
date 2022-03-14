package global

import (
	"log"
	"net"

	"github.com/amitdotkr/sms-otp/src/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOtpServiceServer
}

func GrpcServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server failed to Start %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOtpServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
