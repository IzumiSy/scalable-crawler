.PHONY: deploy run

PROJECT_ID := $$(gcloud config list --format='value(core.project)')

run:
	@PROJECT_ID=$(PROJECT_ID) go run main.go

deploy:
	@gcloud app deploy
