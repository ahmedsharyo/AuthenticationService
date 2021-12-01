package login_grpc

import (
	"context"
	"log"
	"net"

	"github.com/ahmedsharyo/AuthenticationService/controllers"
	"github.com/ahmedsharyo/AuthenticationService/login_grpc/pb"
	"google.golang.org/grpc"
)

func Run() {

	grpcServer := grpc.NewServer()

	var server Server

	pb.RegisterLogingInServer(grpcServer, server)

	listen, err := net.Listen("tcp", "localhost:3001")

	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3001 %v", err)
	}

	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}

// Server is implementation proto interface
type Server struct{}

// Login function responsible to get the Login information
func (Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {

	//Client(request.CameraId)
	controllers.Login(request.Email, request.Password)
	var data pb.LoginResponse
	data.IsUserAuthenticated = true
	return &data, nil
}
