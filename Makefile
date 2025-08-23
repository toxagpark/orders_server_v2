.PHONY: up down psql psql_orders psql_items app-logs redis-logs kafka-logs logs status

up:
	docker-compose up -d

down:
	docker-compose down -v

psql:
	docker-compose exec postgres psql -U postgres -d postgres -c "\dt"

psql_orders:
	docker-compose exec postgres psql -U postgres -d postgres -c "SELECT * FROM orders"

psql_items:
	docker-compose exec postgres psql -U postgres -d postgres -c "SELECT * FROM items"

app-logs:
	docker logs go-app

redis-logs:
	docker-compose logs -f redis

kafka-logs:
	docker-compose logs -f kafka

logs:
	docker-compose logs -f

status:
	docker-compose ps

# При запуске compose автоматическая миграция
# Первой мигрировать ордерс

migration-orders-up:
	docker-compose exec postgres psql -U postgres -d postgres -f /docker-entrypoint-initdb.d/01_create_orders_table.sql

migration-items-up:
	docker-compose exec postgres psql -U postgres -d postgres -f /docker-entrypoint-initdb.d/02_create_items_table.sql

migration-items-down:
	docker-compose exec postgres psql -U postgres -d postgres -f /migrations/down_items_table.sql

migration-orders-down:
	docker-compose exec postgres psql -U postgres -d postgres -f /migrations/down_orders_table.sql