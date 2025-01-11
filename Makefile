postgres:
	# docker run --name some-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17.2-alpine 
	docker-compose -f docker-compose.yml up -d

# createdb
	# docker exec -it postgres17 createdb --username=postgres --owner=postgres go_project_2

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose down 1

dropdb:
	docker exec -it postgres17 dropdb go_project_2

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/redmonkez12/go-project-2/db/sqlc Store

.PHONY: postgres dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock