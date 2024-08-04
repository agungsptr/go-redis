TAG := 1.0
IMAGE := agungsptr/go-redis
CONTAINER := go-redis
COMPOSE := docker compose -f docker-compose.yml


# Redis
build-redis:
	@cp .env build/redis/.env
	docker build -f build/redis/Dockerfile -t $(IMAGE)_redis:7.2.3 build/redis
	@rm build/redis/.env

run-redis:
	@$(COMPOSE) up -d --force-recreate redis

stop-redis:
	@docker container stop $(CONTAINER)_redis || true
	@docker container rm $(CONTAINER)_redis || true


# CRUD Server
build-server:
	docker build -f build/app/Dockerfile -t $(IMAGE)_server:$(TAG) .

run-server:
	@$(COMPOSE) up -d --force-recreate server

stop-server:
	@docker container stop $(CONTAINER)_server || true
	@docker container rm $(CONTAINER)_server || true


# All services (Redis, Server)
build:
	@make -s build-server
	@make -s build-redis

run:
	@$(COMPOSE) down -v || true
	@$(COMPOSE) up -d --force-recreate

stop:
	@$(COMPOSE) down -v || true

purge:
	@make -s stop
	@docker image rm $(IMAGE)_server:$(TAG) || true
	@docker image rm $(IMAGE)_redis:7.2.3|| true


# Others
redis-cli:
	@docker exec -it $(CONTAINER)_redis redis-cli
