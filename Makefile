
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

restart: down up ps


