package main

// running client side Tests in cmd/client/main.go
import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/Soyaka/microlearn-user/api/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}
	fmt.Println("connected")
	defer conn.Close()

	cc := proto.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//RegisterUser(ctx, cc)
	LoginUser(ctx, cc)
	//VerifyToken(ctx, cc)
	//RefreshToken(ctx, cc)

}

func RegisterUser(ctx context.Context, cc proto.UserServiceClient) {

	res, err := cc.RegisterUser(ctx, &proto.RegisterRequest{
		Name:     "hamidi",
		Email:    "hamidi@exassmple.com",
		Password: "password",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

}

func LoginUser(ctx context.Context, cc proto.UserServiceClient) {

	res, err := cc.LoginUser(ctx, &proto.LoginRequest{
		Email:    "hamidi@exassmple.com",
		Password: "password",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}

func VerifyToken(ctx context.Context, cc proto.UserServiceClient) {

	res, err := cc.VerifyToken(ctx, &proto.Token{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImJKaEQ4QGV4YW1wbGUuY29tIiwiTmFtZSI6IlNveWFrYSIsIklEIjoiOWJhNjY4ZmItYWRhMy00YzZjLTkxMDgtY2FkMjg2MjBiZDFmIiwiaXNzIjoidXNlciIsImV4cCI6MTcxMzg4NTM3NX0.Y57Uh2J2vmOimdhvEk-cloaKlqtPe3MuQn9SlF30fFU",
	})

	if err != nil {
		fmt.Println(err)
	}

	log.Println(res)

}

func RefreshToken(ctx context.Context, cc proto.UserServiceClient) {

	res, err := cc.RefreshToken(ctx, &proto.Token{

		Token: "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImJKaEQ4QGV4YW1wbGUuY29tIiwiTmFtZSI6IlNveWFrYSIsIklEIjoiOWJhNjY4ZmItYWRhMy00YzZjLTkxMDgtY2FkMjg2MjBiZDFmIiwiaXNzIjoidXNlciIsImV4cCI6MTcxMzg4NTM3NX0.Y57Uh2J2vmOimdhvEk-cloaKlqtPe3MuQn9SlF30fFU",
	})

	if err != nil {
		fmt.Println(err)
	}

	log.Println(res)

}
