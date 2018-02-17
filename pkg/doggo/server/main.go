package server

import context "golang.org/x/net/context"

const port = ":9000"

func main() {

}

type doggoService struct{}

func (s *doggoService) GetById(ctx context.Context, req *GetByIdRequest) (*Doggo, error) {
	return nil, nil
}

func (s *doggoService) GetAll(*GetAllRequest, DoggoService_GetAllServer) error {
	return nil
}

func (s *doggoService) Save(context.Context, *Doggo) (*Doggo, error) {
	return nil, nil
}

func (s *doggoService) SaveAll(DoggoService_SaveAllServer) error {
	return nil
}
