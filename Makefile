run: postgres createbd migrateup createredis

postgres:
	sudo docker run --name postgres02 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb: 
	sudo docker exec -it postgres02 createdb --username=root --owner=root students
dropdb:
	sudo docker exec -it postgres02 dropdb  students

migrateup: 
	migrate -path migration/ -database "postgresql://root:secret@localhost:5432/students?sslmode=disable" -verbose up

migratedown:
	migrate -path migration/ -database "postgresql://root:secret@localhost:5432/students?sslmode=disable" -verbose down
createredis:
	docker run --name test1 -p 6379:6379 -d redis

proto:
	protoc --go_out=./pkg/student --go_opt=paths=source_relative \
    --go-grpc_out=./api/proto --go-grpc_opt=paths=source_relative \
    api/proto/service.proto


.PHONY: postgres createbd migrateup createredis run
