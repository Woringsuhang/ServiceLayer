package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Woringsuhang/mess/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	flag.Parse()

	// Create tls based credential.
	creds, err := credentials.NewClientTLSFromFile("ca_cert.pem", "x.test.example.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// Set up a connection to the servers.
	conn, err := grpc.Dial("127.0.0.1:6644", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Make a echo client and send an RPC.
	rgc := user.NewUserClient(conn)
	getUser, err := rgc.GetUser(context.Background(), &user.GetUserRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getUser.Data)
}
