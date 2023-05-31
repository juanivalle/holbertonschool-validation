build:
	go build -o awesome-api

run:
	./awesome-api >./awesome-api.log 2>&1 &

stop:
	kill $$(pgrep awesome-api) || true

clean: stop
	rm -f awesome-api awesome-api.log

test:
	@curl http://localhost:9999
	@curl http://localhost:9999/health

help:
	cat Makefile | grep '##' | tr -d '##' | head -n -1
