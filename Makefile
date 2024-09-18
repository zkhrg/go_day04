all: client 

client: docker-req certs server remove-certs
	go build -o bin/client client/main.go
	@echo "пример для запуска ./bin/client -k AA -c 1 -m 50"

server:
	docker compose -f generated-server/docker-compose.yaml up --build -d

certs:
	export PATH=$PATH:/go/bin
	rm -rf certs
	go install github.com/jsha/minica@latest
	mkdir certs
	minica -domains candy.ltd 
	mv candy.ltd certs/candy.ltd
	mv minica-key.pem certs/
	mv minica.pem certs/
	cp -r certs generated-server/certs

docker-req:
	docker pull golang:1.22
	docker pull alpine:latest

remove-certs:
	rm -rf generated-server/certs

clean:
	rm -rf certs 
	rm bin/client
	