.PHONY: build run push
.DEFAULT_GOAL := run

build:
	docker-compose build rc-announcer

run: build
	docker-compose up rc-announcer

push: build
	docker push reschex/rc-announcer:v0.0.1