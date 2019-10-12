package main

import (
	pb "github.com/xiaozefeng/shipper/user-service/proto/user"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Create(c context.Context, u *pb.User, res *pb.Response) error {
	panic("implement me")
}

func (s *service) Get(c context.Context, u *pb.User, res *pb.Response) error {
	panic("implement me")
}

func (s *service) GetAll(c context.Context, req *pb.Request, res *pb.Response) error {
	panic("implement me")
}

func (s *service) Auth(c context.Context, req *pb.User, res *pb.Token) error {
	panic("implement me")
}

func (s *service) ValidateToken(c context.Context, token *pb.Token, res *pb.Token) error {
	panic("implement me")
}
