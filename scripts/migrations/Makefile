ENV_FILE=../../.env

include $(ENV_FILE)

# Автоматически экспортировать переменные из .env
export $(shell sed 's/=.*//' $(ENV_FILE))

DB_URI=$(DB_DRIVER)://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

MIG_DIR=../../deployments/local/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIG_DIR) $(name) # make migrate-create name=create_users

.PHONY: migrate-up
migrate-up:
	@migrate -database $(DB_URI) -path $(MIG_DIR) up

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_URI) -path $(MIG_DIR) down