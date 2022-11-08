.DEFAULT_GOAL := build

.PHONY: build
build:
	go mod tidy
	go mod vendor

.PHONY: image
image:
	docker build -t proxy .
