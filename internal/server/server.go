package server

import (
	"go-microservice/internal/database"
	"log"
	"net/http"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB: db,
	}
	server.RegisterRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080");err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred : %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) RegisterRoutes(){
	
}