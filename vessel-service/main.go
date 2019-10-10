package main

import (
	"errors"
	pb "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

func main() {

}

type Repository interface {
	FindAvailable(specification *pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// Our grpc service handler
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	
}

//func (repo *VesselRepository) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
//}
