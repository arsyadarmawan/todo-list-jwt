DB_URL=postgresql://postgres:admin123@localhost:2022/new-app?sslmode=disable
download:
	go mod download

dep:
	go mod tidy

run:
	go run main.go

testing:
	go test test/tasks_test.go -v

build:
	go build -o bin/moonlay ./main.go

docker-image:
	docker build -t moonlay:v1 .

docker-run:
	docker run -it -d -p 3000:3000 --name moonlay

migrate:
	migrate create -ext sql -dir db/migrations create_table_user

postgres:
	docker run --name postgres -p 2022:5432 -e POSTGRES_PASSWORD=admin123 -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres new-app

dropdb:
	docker exec -it postgres dropdb new-app

migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down



.PHONY: createdb createdb dropdb