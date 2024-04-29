include .env

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: build
build:
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: templ-watch
templ-watch:
	templ generate --watch -proxy=http://localhost:${PORT}

.PHONY: go-watch
go-watch:
	templ generate --watch --cmd="go run ./cmd/$(APP_NAME)/main.go"

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: create-migration
create-migration:
	migrate create -ext=sql -dir=internal/database/migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path=internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up


.PHONY: migrate-down
migrate-down:
	migrate -path=internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

