start:
	echo "Start"

migrate:
	go run cmd/database/migrate.go

seed:
	go run cmd/database/migrate.go seed users

	go run cmd/database/migrate.go seed items