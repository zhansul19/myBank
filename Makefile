DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgresSimple  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgresSimple createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresSimple dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test ./... -cover

server:
	go run main.go

mockgen:
	mockgen -destination db/mock/store.go github.com/zhansul19/myBank/db/sqlc Store

.PHONY:  postgres createdb dropdb migrateup migratedown sqlc test server mockgen