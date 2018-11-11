version = 0.0.1

build:
	CGO_ENABLED=0 go build -o subscriber .

docker:
	docker build -t subscriber:$(version) .

.PHONY: build

