todo:
	go run ./cmd/todo/ --config ./cmd/todo/config/config.toml

swag:
	swag init --generalInfo internal/http/server/server.go --outputTypes yaml --output ./cmd/todo/docs