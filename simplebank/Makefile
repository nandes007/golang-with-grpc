DB_URL=postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres17 --network bank-network -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path=db/migration -database "$(DB_URL)" --verbose up

migrateup1:
	migrate -path=db/migration -database "$(DB_URL)" --verbose up 1

migratedown:
	migrate -path=db/migration -database "$(DB_URL)" --verbose down

migratedown1:
	migrate -path=db/migration -database "$(DB_URL)" --verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

# This need to run because when i run test using wsl it cache the test
cleantestcache:
	go clean -testcache

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nandes007/simplebank/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:8.0-M04-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test cleantestcache server mock proto evans redis