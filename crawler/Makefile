.PHONY: build push

build: Dockerfile
	docker build -t asia.gcr.io/gae_crawler/crawler:latest

push: Dockerfile
	docker push asia.gcr.io/gae_crawler/crawler:latest
