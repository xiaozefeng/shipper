package main

import (
	"encoding/json"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	pb "github.com/xiaozefeng/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cmd.Init()
	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Coulnd not parse file: %v", err)
	}
	response, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created:%v\n", response.Created)

	resp, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignment:%v", err)
	}
	for _, consignment := range resp.Consignments {
		log.Println(consignment)
	}
}

func v1() {
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("failed to connect grpc server: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewShippingServiceClient(conn)
	//
	//// 和服务器通信，并打印返回信息
	//file := defaultFilename
	//if len(os.Args) > 1 {
	//	file = os.Args[1]
	//}
	//
	//consignment, err := parseFile(file)
	//if err != nil {
	//	log.Fatalf("Coulnd not parse file: %v", err)
	//}
	//response, err := client.CreateConsignment(context.Background(), consignment)
	//if err != nil {
	//	log.Fatalf("Could not greet: %v", err)
	//}
	//log.Printf("Created:%v\n", response.Created)
	//
	//resp, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	//if err != nil {
	//	log.Fatalf("Could not list consignment:%v", err)
	//}
	//for _, consignment := range resp.Consignments {
	//	log.Println(consignment)
	//}
}

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	return consignment, err
}
