local:
	@echo "=============starting locally============="
	docker-compose -f resources/docker/docker-compose.yaml up
down:
	docker-compose -f resources/docker/docker-compose.yaml down
down-all:
	@echo "=============Stop All Container============="
	docker stop $$(docker ps -aq)
test:
	export GIN_MODE=release && go test ./app/... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt
lint:
	golangci-lint run --verbose
clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f
	docker images prune -f
format:
	go fmt ./app/...
migrate-create:
	migrate create -ext sql -dir app/migrations $(name)
REPOS_DIR = ./app/repositories/*.go
USE_CASES_DIR = ./app/usecases/*.go
mock: \

	for f in ${REPOS_DIR} ; do \
		mockgen -source=./app/repositories/$$(basename $$f) -destination=./app/mocks/repositories/$$(basename $$f); \
	done

	for f in ${USE_CASES_DIR} ; do \
		mockgen -source=./app/usecases/$$(basename $$f) -destination=./app/mocks/usecases/$$(basename $$f); \
	done
