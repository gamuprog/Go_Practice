# App
.PHONY: run-dev
run-dev:
	docker compose up

.PHONY: destroy
destroy:
	docker-compose down --volumes

# DB
.PHONY: init-db-dev
init-db-dev:
	docker compose -f docker-compose.dev.yml run --rm api go run cmd/initdb/initdb.go dev

## TODO: init周りの作成