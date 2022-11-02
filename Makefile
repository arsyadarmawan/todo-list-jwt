dep:
	go mod tidy

run:
	go run main.go

build:
	go build -o bin/moonlay ./main.go

docker-image:
	docker build -t moonlay:v1 .

docker-run:
	docker run -it -d -p 3000:3000 --name moonlay