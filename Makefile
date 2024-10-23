build:
	go build -o app

test:
	go test ./...

run:
	go run main.go

docker-build:
	docker build -t firefly-app .

docker-run:
	docker run --rm firefly-app
