postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createbd: 
	sudo docker exec -it postgres createdb --username=root --owner=root students
dropdb:
	sudo docker exec -it postgres dropdb  students

migrateup: 
	migrate -path migration/ -database "postgresql://root:secret@localhost:5432/students?sslmode=disable" -verbose up

migratedown:
	migrate -path migration/ -database "postgresql://root:secret@localhost:5432/students?sslmode=disable" -verbose down
.PHONY: postgres createbd dropdb migrateup migratedown
