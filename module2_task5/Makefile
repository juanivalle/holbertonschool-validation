lint: ## lint
	golangci-lint run ./main.go

build: ## Build awesome-api
	make go-build
	make hugo-build

run: ## Run awesome-api application
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Stop awesome-api application
	kill $$(pgrep awesome-api) || true

post: ## Create new MarkDown
	hugo new posts/$(POST_NAME).md
	sed -i 's/title: "$(POST_NAME)"/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

clean: stop ## Clean awesome-api application and log
	rm -f awesome-api awesome-api.log coverage-units.out coverage-integrations.out go*
	rm -rf ./dist

test: ## Test current connection
	make unit-tests
	make integration-tests

unit-tests: ## Run tests to check for HelloHandler coverage
	go test -v  -short -coverprofile=coverage-units.out

integration-tests: ## Run tests to check integrations coverage
	go test -v -coverprofile=coverage-integrations.out

check: ## Check markdown files
	markdownlint ./content/posts/*.md
	markdown-link-check ./content/posts/*.md

validate: ## Validate HTML with W3C Validator
	make build
	python3 /opt/W3C-Validator/w3c_validator.py ./dist/index.html || true

go-build:	## Build API library
	go mod init github.com/holbertonschool-validation/tree/main/module2_task5/awesome-api
	go get github.com/gorilla/mux
	go build

hugo-build: check ## Build go hugo website
	hugo -d dist/

.SILENT:
help: ## Display what each command does
	cat Makefile | grep '##' | tr -d '##' | head -n -1
