package server

import "github.com/toyota-connected/doggo/pkg/doggo/pb"

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
