run:
	go run cmd/main.go

up:
	docker-compose up --build -d

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

db:
	docker exec -it ebd8a0b2bac7 /bin/bash