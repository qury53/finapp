include .env
export
export PROJECT_ROOT=$(shell pwd)

run:
	go run ${PROJECT_ROOT}/cmd/finapp/main.go

docker-up:
	@docker compose up -d 

docker-down:
	@docker compose down

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует параметр seq. make migrate-create seq=название-миграции"; \
		exit 1; \
	fi
	docker compose run --rm finapp-postgres-migrate create -ext sql -dir /migrations -seq "$(seq)"
migrate-up:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm finapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@finapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			up; \
	else \
		docker compose run --rm finapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@finapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			up "$(level)"; \
	fi


migrate-down:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm finapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@finapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			down; \
	else \
		docker compose run --rm finapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@finapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			down "$(level)"; \
	fi

clean-up:
	@read -p "Очистить все volume файлы окружений? Данные будут утеряны. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down finapp-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы очищены"; \
	else \
		echo "Очистка отменена"; \
	fi

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/finapp/main.go \
		-o docs \
		--parseInternal \
		--parseDependency