package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/xiaozefeng/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"log"
)

type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Here we call a client instance of out vessel service with our consignment weight
	// and the amount of containers as the capacity value
	vesselResp, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResp.Vessel.Name)
	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResp.Vessel.Id

	// save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

//func (s *service) GetConsignments(context.Context, *pb.GetRequest) (*pb.Response, error) {
//	consignments := s.repo.GetAll()
//	return &pb.Response{Consignments: consignments}, nil
//
//}
//
//func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
//	consignment, err := s.repo.Create(req)
//	if err != nil {
//		return nil, err
//	}
//	return &pb.Response{Created: true, Consignment: consignment}, nil
//}

func main() {
	repo := &ConsignmentRepository{}

	// Create a new service.
	// Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// 	Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo: repo, vesselClient: vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}

func v1() {
	//repo := &Repository{}
	//
	//// 启动 rpc 服务
	//listener, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//s := grpc.NewServer()
	//// 注册服务到 gRPC 服务器, 会把自定义的 protobuf 与自动生成的代码接口进行绑定
	//pb.RegisterShippingServiceServer(s, &service{repo})
	//
	//// 在 gRPC 服务器上注册 reflection 服务
	//reflection.Register(s)
	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
