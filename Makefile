
up:
	docker-compose up -d

down:
	docker-compose down

ps:
	docker-compose ps

bash:
	docker-compose run --rm go bash

restart: down up ps


