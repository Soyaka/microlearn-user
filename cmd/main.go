package main

import (
	"log"
	"net"
	"os"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"github.com/Soyaka/microlearn-user/internal/handlers"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	var port string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		port = "50052"
	}
	port = os.Getenv("GRPC_PORT")
	Addr := ":" + port

	listener, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	UserService := handlers.NewUmplimentUserMethods()
	RegisterServerService(server, UserService)
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func RegisterServerService(server *grpc.Server, service *handlers.UmplimentUserMethods) {
	proto.RegisterUserServiceServer(server, service)
}
