package procedure

import (
	"errors"
	"log"

	"github.com/RajibDas-123/ms-grpc-auth/auth/pb"
	"golang.org/x/net/context"
	// replace this with your own project
)

// AuthServer : Server
type AuthServer struct {
}

// Login : login procedure
func (s *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	email := request.GetEmailId()
	pass := request.GetPassword()
	log.Println("Authenticating ", email, pass)

	var response pb.LoginResponse

	if email == "raj@adl.com" && pass == "password" {
		response.Success = true
		return &response, nil
	}
	log.Println("Login request!!")
	response.Success = false
	return &response, errors.New("login: unable to login")
}

// Logout : logout procedure
func (s *AuthServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}

// IsSessionExist : isSessionExistn procedure
func (s *AuthServer) IsSessionExist(ctx context.Context, request *pb.IsSessionExistRequest) (*pb.IsSessionExistResponse, error) {
	return nil, nil
}

// GetSessionData : getSessionData procedure
func (s *AuthServer) GetSessionData(ctx context.Context, request *pb.GetSessionDataRequest) (*pb.GetSessionDataResponse, error) {
	return nil, nil
}

// SetSessionData : setSessionData procedure
func (s *AuthServer) SetSessionData(ctx context.Context, request *pb.SetSessionDataRequest) (*pb.SetSessionDataResponse, error) {
	return nil, nil
}
