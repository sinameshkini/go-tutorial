tidy:
	go mod tidy

run:
	go run cmd/main.go

build:
	go build -o build/auth cmd/main.go

up:
	docker-compose up -d

down:
	docker-compose down