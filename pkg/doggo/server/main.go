package server

import (
	"github.com/toyota-connected/doggo/pkg/doggo/pb"
	context "golang.org/x/net/context"
)

const port = ":9000"

func main() {

}

type doggoService struct{}

func (s *doggoService) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Doggo, error) {
	return nil, nil
}

func (s *doggoService) GetAll(*pb.GetAllRequest, pb.DoggoService_GetAllServer) error {
	return nil
}

func (s *doggoService) Save(context.Context, *pb.Doggo) (*pb.Doggo, error) {
	return nil, nil
}

func (s *doggoService) SaveAll(pb.DoggoService_SaveAllServer) error {
	return nil
}
