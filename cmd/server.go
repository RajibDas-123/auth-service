package cmd

import (
	"net"

	"github.com/RajibDas-123/ms-grpc-auth/auth/logging"
	"github.com/RajibDas-123/ms-grpc-auth/auth/pb"
	"github.com/RajibDas-123/ms-grpc-auth/auth/procedure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run : runs the server
func Run() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		logging.AppLogger.Fatalf("Failed to listen:  %v", err)
	}
	logging.AppLogger.Infof("listening on :3000 ")
	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, &procedure.AuthServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logging.AppLogger.Fatalf("Failed to serve: %v", err)
	}
}
