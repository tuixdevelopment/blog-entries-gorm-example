-include .env

build: clean
	go build -o ./bin/stellar ./api/server/main.go

start: build
	docker-compose up -d && sleep 5
	go run ./api/server/main.go

stop:
	docker-compose down

clean:
	rm -rf ./bin
	rm -rf bootstrap