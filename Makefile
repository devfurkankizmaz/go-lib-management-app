migration:
	go run migrate/migrate.go

app:
	go run main.go

dev:
	docker compose up -d

dev-down:
	docker compose down


.PHONY: dev dev-down app migration