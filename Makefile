
hello:
	@echo "Hello world"


migrateup:
	 migrate -path db/migrations -database "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrateup1:
	 migrate -path db/migrations -database "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migratedown:
	 migrate -path db/migrations -database "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable" --verbose down

migratedown1:
	 migrate -path db/migrations -database "postgresql://ibrahim:ibrahim@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

proto:
	rm -f pb/*.go
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
run:
	go run main.go

sqlc: 
	sqlc generate
	
server:
	go run main.go

test:
	go test -v -cover ./...

mock:
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
	mockgen -package mockdb  -destination db/mock/store.go  github.com/baksman/backend_masterclass/db/sqlc Store



.PHONY: proto  sqlc migratedown migrateup test mock migrateup1 migratedown1