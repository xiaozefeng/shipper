build:
	protoc -I. --go_out=plugins=micro:$(HOME)/github.com/go/shipper/consignment-service proto/consignment/*.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 \
	 -e micro_server_address=:50051 \
	 -e micro_registry=mdns consignment-service

clean:
	rm -rf consignment-service