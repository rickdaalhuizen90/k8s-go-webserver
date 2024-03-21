.PHONY: dev prod build

dev:
	skaffold dev -p dev

prod:
	skaffold build -p prod

build:
	docker build -t go-webserver:latest ./src
