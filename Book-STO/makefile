proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative test_microservice/proto/services.proto

build:
	docker compose build --no-cache

run:
	docker compose up