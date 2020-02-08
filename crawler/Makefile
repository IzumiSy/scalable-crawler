.PHONY: build push

PROJECT_ID := $$(gcloud config list --format='value(core.project)')
IMAGE_PATH := asia.gcr.io/$(PROJECT_ID)/crawler:latest

build: Dockerfile
	@docker build -t crawler .

push: Dockerfile
	@docker tag crawler $(IMAGE_PATH)
	@docker push $(IMAGE_PATH)
