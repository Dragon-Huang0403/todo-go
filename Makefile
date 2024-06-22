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
