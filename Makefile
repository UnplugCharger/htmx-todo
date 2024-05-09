DB_NAME=htmx_to_do_db
DB_URI=postgresql://root:password@localhost:5432/$(DB_NAME)?sslmode=disable
DB_MIGRATIONS_PATH=./db/migrations
DB_SEEDS_PATH=./db/seeds
DB_USER=root
DB_PASSWORD=password
DB_USER2=postgres

postgres:
	docker run --name datapoint_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it datapoint_db createdb --username=$(DB_USER) --owner=root $(DB_NAME)
dropdb:
	docker exec -it datapoint_db dropdb $(DB_NAME)

migrateup:
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_URI) --verbose up
migrateup1:
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_URI) --verbose up 1

migratedown:
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_URI) --verbose down
migratedown1:
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DB_URI) --verbose down 1

migratecreate:
	migrate create -ext sql -dir $(DB_MIGRATIONS_PATH) -seq $(name)

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

run:
	air -c .air.toml