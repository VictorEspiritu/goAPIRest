
up:build
	docker-compose up -d

down:
	docker-compose down

ps:
	docker-compose ps

bash:
	docker-compose exec go bash

build:
	docker-compose build 
	
logs:
	docker-compose logs -f

test: test-integration

test-integration:
	@docker-compose exec go go test -count=1 ./test/...  

restart: down up ps


