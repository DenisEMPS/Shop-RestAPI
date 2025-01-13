APP_NAME=app

.PHONY: build, run, docker_postgres, create_migration, migrate, migrate_down, test_client, test_supplier, gen

docker_postgres:
	docker run --name=shop_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

build:
	go build -o $(APP_NAME) cmd/main.go

run: docker_postgres build
	$(APP_NAME)

create_migration:
	migrate create -ext sql -dir schema/ -seq $(NAME)

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

gen:
	mockgen -source=/pkg/service/service.go -destination=/pkg/service/mock.go

test_client:
	go test -count=1 ./test/client

test_supplier:
	go test -count=1 ./test/supplier



	