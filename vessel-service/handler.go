package main

import (
	pb "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
)

// rpc handler
type service struct {
	session *mgo.Session
}

func (s *service) FindAvailable(c context.Context, req *pb.Specification, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	// find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

func (s *service) Create(c context.Context, v *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(v); err != nil {
		return err
	}
	res.Vessel = v
	res.Created = true
	return nil
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{session: s.session.Clone()}
}
