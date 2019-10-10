build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run -e micro_registry=mdns consignment-cli

clean:
	rm -rf consignment-cli