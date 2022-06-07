.PHONY: run
## Run service. Usage: 'make run'
run: ; $(info Running...)
	go run main.go

.PHONY: test
## Run tests. Usage: 'make test'
test: ; $(info running tests…)
	docker-compose -f ./docker-compose.test.yml up --force-recreate -d; \
	go test -v -failfast ./...
	docker-compose -f ./docker-compose.test.yml down;
