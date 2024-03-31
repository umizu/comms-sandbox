package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

var jobs = make(map[string]*Job)

func main() {
	server := NewAPIServer(":8080")
	server.Start()
}

func (s *APIServer) Start() {
	s.router.POST("/submit", submitHandler)
	s.router.GET("/checkstatus", checkStatusHandler)

	s.router.Logger.Fatal(s.router.Start(s.listenAddr))
}

type APIServer struct {
	router     *echo.Echo
	listenAddr string
}

type Job struct {
	Progress int
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		router:     echo.New(),
		listenAddr: listenAddr,
	}
}

func submitHandler(c echo.Context) error {
	id := random.String(7) // weak
	jobs[id] = &Job{
		Progress: 0,
	}
	go updateJob(id)
	return c.String(http.StatusOK, fmt.Sprintf("job:%s", id))
}

func checkStatusHandler(c echo.Context) error {
	id := c.QueryParam("jobId")
	job := jobs[id]
	if job == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("jobId %s not found", id))
	}

	return c.String(http.StatusOK, fmt.Sprintf("status:%d%%", job.Progress))
}

func updateJob(id string) {
	for jobs[id].Progress < 100 {
		time.Sleep(5 * time.Second)
		jobs[id].Progress += 10
	}
}
