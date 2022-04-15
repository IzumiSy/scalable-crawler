# Scalable Crawler
CloudNative scalable crawler built on Google Cloud Platform

![architecture](docs/architecture.jpeg)

## Why Cloudbuild?
If we just want to run crawler application, there are options such as GAE, Cloud Run, and more, but CloudBuild has the way longer timeout. It can runs application to the extent of 24 hours long at maximum. The longer the better because crawling is the process that sometimes takes long time.

## Setup

### Crawler
```bash
$ make build
$ make push
```

### Backend
```bash
$ make deploy
```

