GIT_SHA:=$(shell git rev-parse HEAD)
build:
	docker build -t bencord0/hello:$(GIT_SHA) .

push:
	docker push bencord0/hello:$(GIT_SHA)
