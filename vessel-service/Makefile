generate:
	protoc -I. --go_out=plugins=micro:$(HOME)/github.com/go/shipper/vessel-service proto/vessel/*.proto

build:generate
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns vessel-service

clean:
	rm -f vessel-service
