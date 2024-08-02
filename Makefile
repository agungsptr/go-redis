TAG := 1.0
IMAGE := agungsptr/go-redis
CONTAINER := go-redis
COMPOSE := docker compose -f docker-compose.yml


# CRUD Server
build-server:
	docker build -t $(IMAGE)_server:$(TAG) .

run-server:
	@$(COMPOSE) up -d --force-recreate server

stop-server:
	@docker container stop $(CONTAINER)_server || true
	@docker container rm $(CONTAINER)_server || true


# All services (Redis, Server)
build:
	@make -s build-server

run-services:
	@$(COMPOSE) down -v || true
	@$(COMPOSE) up -d --force-recreate

stop-services:
	@$(COMPOSE) down -v || true

purge-services:
	@make -s stop-services
	@docker image rm $(IMAGE)_server:$(TAG) || true


# Others
redis-cli:
	@docker exec -it $(CONTAINER)_redis redis-cli
