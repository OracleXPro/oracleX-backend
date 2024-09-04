#include .env

NETWORK=oraclex 
PROJECT=oraclex 


.PHONY: goZero
goZero:
	goctl rpc protoc -m --proto_path=api/proto api/proto/*.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style goZero

.PHONY: swaggerJson
swaggerJson:
	protoc --proto_path=api/proto api/proto/*.proto --openapiv2_out=./api/swagger

.PHONY: validation
validation:
	protoc --proto_path=api/proto api/proto/*.proto  --validate_out=paths=source_relative,lang=go:./api/pb

.PHONY: gateway
gateway:
	protoc --include_imports --proto_path=./api/proto --descriptor_set_out=./api/pb/descriptor.pb oracleX.proto

.PHONY: merge
merge:
	cd api/proto && sh merge.sh

.PHONY: gen
gen: merge goZero validation gateway swaggerJson

.PHONY: swagger
swagger:
	docker run -p 8082:8080 -e SWAGGER_JSON=/swagger/swagger.json  -v ./api/swagger/monozero.swagger.json:/swagger/swagger.json swaggerapi/swagger-ui:v5.10.3

.PHONY: db
db:
	goctl model mysql datasource --url="root:letmein@tcp(127.0.0.1:33060)/oracle_x" --table="user,telegram,wallet,wallet_login_nonce" -dir ./internal/model/db

.PHONY: model
model: db

.PHONY: run
run:
	go run oracleX.go

.PHONY: build
build:
	DOCKER_BUILDKIT=1  BUILDKIT_PROGRESS=plain docker build -t app -f Dockerfile .

.PHONY: healthy
healthy:
	curl localhost:8888/ping/health


.PHONY: network
network:
	docker network create $(NETWORK)

.PHONY: up
up:
	cd deployment && docker-compose --env-file ../.env -p $(PROJECT) -f docker-compose.yaml up -d


.PHONY: down
down:
	cd deployment && docker-compose --env-file ../.env -p $(PROJECT) -f docker-compose.yaml down
