package main

import (
	"github.com/micro/go-micro"
	pb "github.com/xiaozefeng/shipper/vessel-service/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}
	repo := &VesselRepository{session: session.Copy()}
	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"))
	srv.Init()

	// Register out implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session: session})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "jack", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		_ = repo.Create(v)
	}
}
