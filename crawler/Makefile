.PHONY: build push

build: Dockerfile
	docker build -t crawler .

push: Dockerfile
	docker tag crawler asia.gcr.io/gae_crawler/crawler:latest
	docker push asia.gcr.io/gae_crawler/crawler:latest
