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
SERVER_GREETING=todo-greeting
SERVER_TODO=todo-todo
SERVER_USER=todo-user
SERVER_APPMIXER=todo-appmixer
into-greeting:
	cd .docker && docker compose exec ${SERVER_GREETING} /bin/sh
into-todo:
	cd .docker && docker compose exec ${SERVER_TODO} /bin/sh
into-user:
	cd .docker && docker compose exec ${SERVER_USER} /bin/sh
into-appmixer:
	cd .docker && docker compose exec ${SERVER_APPMIXER} /bin/sh

# run grpc on foregrond
run-greeting:
	cd .docker && docker compose exec ${SERVER_GREETING} /bin/sh -c "go run /go/src/app/services/greeting/main.go"
run-todo:
	cd .docker && docker compose exec ${SERVER_TODO} /bin/sh -c "go run /go/src/app/services/todo/main.go"
run-user:
	cd .docker && docker compose exec ${SERVER_USER} /bin/sh -c "go run /go/src/app/services/user/main.go"
run-appmixer:
	cd .docker && docker compose exec ${SERVER_APPMIXER} /bin/sh -c "go run /go/src/app/services/appmixer/main.go"

# run grpc on background
run-greeting-back:
	cd .docker && docker compose exec -d ${SERVER_GREETING} /bin/sh -c "go run /go/src/app/services/greeting/main.go"
run-todo-back:
	cd .docker && docker compose exec -d ${SERVER_TODO} /bin/sh -c "go run /go/src/app/services/todo/main.go"
run-user-back:
	cd .docker && docker compose exec -d ${SERVER_USER} /bin/sh -c "go run /go/src/app/services/user/main.go"
run-appmixer-back:
	cd .docker && docker compose exec -d ${SERVER_APPMIXER} /bin/sh -c "go run /go/src/app/services/appmixer/main.go"
run-all-back:
	./start-grpc.sh

# proto builder
BUILD_FROM=proto/services
proto-gen:
	protoc --go_out=. \
        --go-grpc_out=. \
        ${BUILD_FROM}/*

# grpc-gateway builder
proto-gw:
	protoc -I ./proto \
		--go_out=./ \
		--go-grpc_out=. \
		--grpc-gateway_out=. \
		--plugin=protoc-gen-grpc-gateway=${GOBIN}/protoc-gen-grpc-gateway \
		proto/gateway/appmixer.proto

# Check each grpc server working or no
ls-grpc-greeting:
	grpc_cli ls localhost:4000
ls-grpc-todo:
	grpc_cli ls localhost:4001
ls-grpc-user:
	grpc_cli ls localhost:4002

# Check each grpc server working on first time
# With updating day by day, this funcions are disable
call-grpc-greeting:
	grpc_cli call localhost:4000 greeting.Greeter.SayHello 'name:"Twice!!!!!"'
call-grpc-todo:
	grpc_cli call localhost:4001 todo.Greeter.SayHello 'name:"Twice!!!!!"'
call-grpc-user:
	grpc_cli call localhost:4002 user.Greeter.SayHello 'name:"Twice!!!!!"'

# Check grpc-gateway running
sample-gw:
	curl -X POST -k http://localhost:8000/todo/say-hello -d '{"name": "Hinako!!!!"}'