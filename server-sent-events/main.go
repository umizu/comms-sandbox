package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/stream", streamHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func streamHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	return send(c, 5)
}

func send(c echo.Context, count int) error {
	if count == 0 {
		return nil
	}

	c.Response().Write([]byte(
		fmt.Sprintf("data: hello from server (%d)\n\n", count)))
	time.Sleep(time.Second) // simulate long operation
	return send(c, count-1)
}
