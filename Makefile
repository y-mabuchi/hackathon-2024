.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: stop
stop:
	docker compose stop

.PHONY: restart
restart:
	docker compose restart

.PHONY: build
build:
	docker compose up --build

.PHONY: protoc
protoc:
	cd proto && \
	protoc \
	--go_out=../backend/handler/grpc/generated --go_opt=paths=source_relative \
	--go-grpc_out=../backend/handler/grpc/generated --go-grpc_opt=paths=source_relative \
	*.proto
