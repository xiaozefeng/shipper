package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/xiaozefeng/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
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
	repo := &Repository{}

	// Create a new service.
	// Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 	Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

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
