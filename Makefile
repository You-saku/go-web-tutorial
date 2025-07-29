setup:
	docker compose up -d --build
seeding:
	docker compose exec -T db mysql --user=user --password=secret develop < database/migrations/create-users-table.sql
	sleep 1
	docker compose exec -T db mysql --user=user --password=secret develop < database/seed/insert-user.sql
up:
	docker compose up -d
down:
	docker compose down
ps:
	docker compose ps
go:
	docker compose exec web sh
db:
	docker compose exec db mysql --user=user --password=secret
start:
	docker compose exec web go run main.go
