package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/cloudbuild/v1"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		panic("projectId is required")
	}

	e := echo.New()
	http.Handle("/", e)
	e.GET("/", rootHandler())
	e.POST("/queue", queueHandler(projectId))

	port := 8080
	if envPort := os.Getenv("PORT"); envPort != "" {
		_envPort, err := strconv.Atoi(envPort)
		if err != nil {
			panic(err)
		}
		port = _envPort
	}

	log.Printf("Listening on port: %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func rootHandler() func(e echo.Context) error {
	return func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello")
	}
}

func queueHandler(projectId string) func(e echo.Context) error {
	return func(e echo.Context) error {
		ctx := context.Background()

		cb, err := cloudbuild.NewService(ctx)
		if err != nil {
			return e.String(http.StatusInternalServerError, "failure")
		}

		crawl := cloudbuild.BuildStep{
			Name: fmt.Sprintf("asia.gcr.io/%s/crawler:latest", projectId),
			Args: []string{"https://google.com"},
		}
		params := cloudbuild.Build{
			Steps: []*cloudbuild.BuildStep{&crawl},
		}

		build := cloudbuild.NewProjectsBuildsService(cb)
		operation, err := build.Create(projectId, &params).Do()
		if err != nil {
			log.Printf("Failed: %s", err.Error())
			return e.String(http.StatusInternalServerError, "failure")
		}

		return e.String(http.StatusCreated, operation.Name)
	}
}
