package main

import (
	"github.com/micro/go-micro"
	pb "github.com/xiaozefeng/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// 从环境变量中获取 database host
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Could not connect to database with host %s-%v", host, err)
	}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"))
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session: session, vesselClient: vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
