#Create Container
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

# Create DB
createdb :
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

# Drop DB
dropdb :
	docker exec -it postgres12 dropdb  simplebank

# Migrate the DB up
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

# Migrate the DB down
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc :
	sqlc generate

test : 
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb test sqlc migrateup migratedown