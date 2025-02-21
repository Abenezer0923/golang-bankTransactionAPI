.PHONY: postgres createdb dropdb migrateup migratedown sqlc test postgres-connect show-db mock

# Start PostgreSQL container
postgres:
	sudo docker-compose up -d

# Show running containers
show:
	sudo docker ps 

# Create the database
createdb:
	docker exec -it bank_postgres createdb --username=root --owner=root simple_bank

# Drop the database
dropdb:
	docker exec -it bank_postgres dropdb simple_bank

# Run migrations
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

# Generate code using sqlc
sqlc:
	sqlc generate

# Connect to PostgreSQL inside Docker
postgres-connect:
	sudo docker exec -it bank_postgres psql -U root -d simple_bank

# Show all tables in the database
show-db:
	sudo docker exec -it bank_postgres psql -U root -d simple_bank -c "\dt"

# Show all data from a specific table (replace 'accounts' with your table name)
show-accounts:
	sudo docker exec -it bank_postgres psql -U root -d simple_bank -c "SELECT * FROM accounts;"

# Generate mocks
mock:
	 mockgen -package mockdb -destination db/mock/store.go github.com/Abenezer0923/simple-bank/db/sqlc Store

server:
	go run main.go


# Run tests
test:
	go test -v ./...