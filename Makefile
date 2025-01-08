postgres:
	# docker run --name some-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17.2-alpine 
	docker-compose -f docker-compose.yml up -d

# createdb
	# docker exec -it postgres17 createdb --username=postgres --owner=postgres go_project_2

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_project_2?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres17 dropdb go_project_2

sqlc:
	sqlc generate

.PHONY: postgres dropdb migrateup migratedown sqlc