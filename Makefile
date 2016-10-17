TARGETS := docker-build fmt
.PHONY: $(TARGETS)
.DEFAULT_GOAL := build/bakery-agent

build/bakery-agent: build/.build-image.log
	docker run --rm -v $(PWD):/usr/src/bakery-agent bakery-agent go build -o $@

build:
	mkdir ./build

build/.build-image.log: build
	docker build -t bakery-agent ./ &> $@

fmt:
	docker run --rm -v $(PWD):/usr/src/bakery-agent bakery-agent go fmt

clean:
	rm -rf ./build
	docker rmi $(shell docker images -q bakery-agent:latest)
