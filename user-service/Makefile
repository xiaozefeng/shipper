generate:
	protoc -I. --go_out=plugins=micro:$(HOME)/github.com/go/shipper/user-service proto/user/*.proto

build:generate
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
	docker run -p 50053:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns user-service

clean:
	rm -f vessel-service