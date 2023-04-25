api:
	go run .\cmd\api\main.go

postgres:
	docker run --name postgres15.2a -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=loladmit -d postgres:15.2-alpine

postgrestty:
	docker exec -it postgres15.2a psql

createdb:
	docker exec -it postgres15.2a createdb --username=root --owner=root pet_feeder

dropdb:
	docker exec -it postgres15.2a dropdb pet_feeder

migrateup:
	migrate -path internal/database/migrations -database "postgresql://root:loladmit@localhost:5432/pet_feeder?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/database/migrations -database "postgresql://root:loladmit@localhost:5432/pet_feeder?sslmode=disable" -verbose down

.PHONY: api