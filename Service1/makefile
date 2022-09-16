proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative test_microservice/proto/services.proto

build:
	docker build -t service1 .

run:
	docker run --name service1 -p 4040:4040 service1

connect:
	docker network connect book-sto_book-store-net service1

###Go to Docker desktop and reboot container service1