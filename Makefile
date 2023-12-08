

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
# mysql -uroot -p

redis:
	@docker exec -it redis bash
# redis-cli


.PHONY: run up down net