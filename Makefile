build:
	go build -o app

test:
	go test ./...

run:
	go run main.go

docker-build:
	docker build -t wordcounterproject .

docker-run:
	docker run --rm wordcounterproject
