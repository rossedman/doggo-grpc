package main

import (
	"context"
	"flag"
	"log"

	"github.com/rossedman/doggo/pkg/doggo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {
	option := flag.Int("o", 1, "Find doggo by id")
	flag.Parse()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewDoggoServiceClient(conn)
	switch *option {
	case 1:
		SendMetadata(client)
	}
}

func SendMetadata(client pb.DoggoServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"rossedman"}
	md["password"] = []string{"password1"}

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	d, err := client.GetById(ctx, &pb.GetByIdRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(d)
}
