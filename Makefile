

run:
	go run ./cmd/*go

up:
	docker-compose up -d

down:
	docker-compose down

net:
	@docker inspect mysql8 | grep IPAddress

mysql:
	@docker exec -it mysql8 bash

.PHONY: run up down net