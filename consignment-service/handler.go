package main

import (
	pb "github.com/xiaozefeng/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	// 使用consignment weight 和容器数量作为Capacity value 生成一个客户端实例
	vesselResp, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel:%v", vesselResp.Vessel.Name)

	// 将从 vessel service 获得的 id 设置为 vesselId
	req.VesselId = vesselResp.Vessel.Id

	// save consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest,resp *pb.Response) error {
	repo:= s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}
