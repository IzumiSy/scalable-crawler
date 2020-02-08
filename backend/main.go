package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	e := echo.New()
	http.Handle("/", e)

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello")
	})

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
