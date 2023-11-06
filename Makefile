postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15.3-alpine

createdb:
	docker exec -it postgres15  createdb --username=root --owner=root bank_app

dropdb:
	docker exec -it postgres15  dropdb bank_app

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_app?sslmode=disable" -verbose down

migrateupwin:
	C:\Users\Alexey\Documents\GoProjects\migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_app?sslmode=disable" -verbose up

migratedownwin:
	C:\Users\Alexey\Documents\GoProjects\migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank_app?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	C:\Users\Alexey\Documents\goworkspace\bin\mockgen -package mockdb -destination db/sqlc/mock/store.go bank-app/db/sqlc Store  

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server