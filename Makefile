todo:
	go run ./cmd/todo/ --config ./cmd/todo/config/config.toml

swag:
	swag fmt
	go fmt ./...
	swag init --generalInfo internal/http/server/server.go --outputTypes yaml --output ./cmd/todo/docs

mock:
	mockgen -destination ./internal/controller/mock/controller.go github.com/dragon-huang0403/todo-go/internal/controller Task
	mockgen -destination ./internal/db/mock/db.go github.com/dragon-huang0403/todo-go/internal/db Database
	mockgen -destination ./internal/store/mock/store.go github.com/dragon-huang0403/todo-go/internal/store Store

build:
	docker buildx bake --set "*.platform=linux/arm64" -f docker-bake.hcl

lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.59.1-alpine golangci-lint run -v