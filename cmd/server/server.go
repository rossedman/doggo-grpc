package main

import (
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/rossedman/doggo-grpc/pkg/doggo/pb"
	context "golang.org/x/net/context"
)

const port = ":9000"

var doggos = []pb.Doggo{
	pb.Doggo{
		Id:     1,
		Name:   "Bongo",
		Breed:  "Shih Tzu",
		IsGood: false,
	},
	pb.Doggo{
		Id:     2,
		Name:   "Oreo",
		Breed:  "Shih Tzu",
		IsGood: true,
	},
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterDoggoServiceServer(s, new(doggoService))
	log.Println("Starting server on port " + port)
	s.Serve(lis)
}

type doggoService struct{}

func (s *doggoService) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Doggo, error) {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("Metadata: %v\n", md)
	}

	log.Print("Request received")

	for _, d := range doggos {
		if req.Id == d.Id {
			return &d, nil
		}
	}

	return nil, errors.New("Doggo not found")
}
