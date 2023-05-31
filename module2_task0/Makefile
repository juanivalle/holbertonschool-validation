build: ## Build awesome-api
	go build -o awesome-api

run: ## Run awesome-api application
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop awesome-api application
	kill $$(pgrep awesome-api) || true

clean: stop ## Clean awesome-api application and log
	rm -f awesome-api awesome-api.log

test: ## Test current connection
	@curl http://localhost:9999
	@curl http://localhost:9999/health

help: ## Display what each command does
	cat Makefile | grep '##' | tr -d '##' | head -n -1
