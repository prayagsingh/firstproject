.PHONY: all dev clean build env-up env-down run

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd firstproject-network && docker-compose up --force-recreate -d
	@echo "Docker environment up"

env-down:
	@echo "Stop environment ..."
	@cd firstproject-network && docker-compose down
	@echo "Docker environment down"

##### RUN
run:
	@echo "Start app and initializing skd with local network..."
	@./firstproject

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/firstproject*
	@docker rm -f $(docker ps -a -q) 2>/dev/null || true
	@echo "Clean up done"