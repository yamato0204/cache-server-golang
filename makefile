
.PHONY: up
up:
	docker compose up --build --watch

test:
	docker compose -f ./docker-compose.yml exec -T go sh -c "cd /app && go test -v -cover ./tests/integration"


db-bash:
	docker compose exec -it db bash


# ディレクトリのパスを変数として定義


# マイグレーション名のデフォルト値
MIGRATION_NAME?=default_migration

#make create-migration MIGRATION_NAME=create_users_table
create-migration:
	migrate create -ext sql -dir api/migrations -seq $(MIGRATION_NAME)
