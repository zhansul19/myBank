DB_URL=postgresql://root:Secretpassword@mybank.cztyee9e1dsx.ap-southeast-1.rds.amazonaws.com:5432/simple_bank

postgres:
	docker run --name postgresSimple  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgresSimple createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresSimple dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1	

sqlc:
	sqlc generate

test:
	go test ./... -cover

server:
	go run main.go

mockgen:
	mockgen -destination db/mock/store.go github.com/zhansul19/myBank/db/sqlc Store

.PHONY:  postgres createdb dropdb migrateup migratedown migratedown1 sqlc test server mockgen