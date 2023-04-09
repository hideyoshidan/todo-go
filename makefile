SERVER_GREETING=todo-greeting
SERVER_TODO=todo-todo
SERVER_USER=todo-user
SERVER_APPMIXER=todo-appmixer
build:
	cd .docker && docker compose build --no-cache
up:
	cd .docker && docker compose up -d
ps:
	cd .docker && docker compose ps
restart:
	cd .docker && docker compose restart
start:
	cd .docker && docker compose start
stop:
	cd .docker && docker compose stop
down:
	cd .docker && docker compose down

# exec into container
into-greeting:
	cd .docker && docker compose exec ${SERVER_GREETING} /bin/sh
into-todo:
	cd .docker && docker compose exec ${SERVER_TODO} /bin/sh
into-user:
	cd .docker && docker compose exec ${SERVER_USER} /bin/sh
into-appmixer:
	cd .docker && docker compose exec ${SERVER_APPMIXER} /bin/sh

# run grpc
run-greeting:
	cd .docker && docker compose exec -d ${SERVER_GREETING} /bin/sh -c "go run /go/src/app/services/greeting/main.go"
run-todo:
	cd .docker && docker compose exec -d ${SERVER_TODO} /bin/sh -c "go run /go/src/app/services/todo/main.go"
run-user:
	cd .docker && docker compose exec -d ${SERVER_USER} /bin/sh -c "go run /go/src/app/services/user/main.go"
run-appmixer:
	cd .docker && docker compose exec -d ${SERVER_APPMIXER} /bin/sh -c "go run /go/src/app/services/appmixer/main.go"
run-all:
	./start-grpc.sh

# Samples
ls-grpc:
	grpc_cli ls localhost:4001
call-grpc:
	grpc_cli call localhost:4001 todo.Greeter.SayHello 'name:"Hinako"'

# protoのコマンド
# https://qiita.com/maaaashin324/items/b8d3c5c016203dce2d6a
BUILD_FROM=src/protofiles/services/
proto-gen:
	protoc --go_out=. \
        --go-grpc_out=. \
        ${BUILD_FROM}/*

proto-gw:
	protoc -I ./proto \
		--go_out=./ \
		--go-grpc_out=. \
		--grpc-gateway_out=. \
		proto/services/appmixer.proto

sample-curl:
	curl -X POST -k http://localhost:8000/todo/say-hello -d '{"name": "Hinako!!!!"}'