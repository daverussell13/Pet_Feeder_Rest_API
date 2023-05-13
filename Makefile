server:
	go run .\cmd\server\main.go

postgres:
	docker run --name postgres15.2a -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=loladmit -d postgres:15.2-alpine

postgrestty:
	docker exec -it postgres15.2a psql

createdb:
	docker exec -it postgres15.2a createdb --username=root --owner=root pet_feeder

dropdb:
	docker exec -it postgres15.2a dropdb pet_feeder

migrateup:
	migrate -path infrastructures/database/migrations -database "postgresql://root:loladmit@localhost:5432/pet_feeder?sslmode=disable" -verbose up

migratedown:
	migrate -path infrastructures/database/migrations -database "postgresql://root:loladmit@localhost:5432/pet_feeder?sslmode=disable" -verbose down -all

dbseed:
	go run .\cmd\seeder\main.go -database "postgresql://root:loladmit@localhost:5432/pet_feeder?sslmode=disable"

migratefresh:
	make migratedown && make migrateup && make dbseed


.PHONY: server migratedown migrateup dropdb createdb postgrestty postgres dbseed migratefresh