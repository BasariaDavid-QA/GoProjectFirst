# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:Vasaner12@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down
	
# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер

gen:
	oapi-codegen -config openapi/.openapi -include-tags messages -package messages openapi/openapi.yaml > ./internal/web/messages/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

generate-users:
	oapi-codegen -generate types,server,spec -o internal/web2/users/users.go -package handlers openapi/openapi.yaml
