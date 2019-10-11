package main

import (
	"errors"
	"github.com/micro/go-micro"
	pb "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"log"
)

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "jack's salty secret", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}
	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	srv.Init()

	// Register out implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to run server:%v", err)
	}
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
	return nil, errors.New("no vessel found by that spec")
}

// Our grpc service handler
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	// find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

//func (repo *VesselRepository) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
//}
