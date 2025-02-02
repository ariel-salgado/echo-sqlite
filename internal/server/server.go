package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	addr string
	echo *echo.Echo
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, echo: echo.New()}
}

func (s *Server) Start() error {
	server:= &http.Server{
		Addr: s.addr,
	}

	go func() {
		if err := s.echo.StartServer(server); err != nil {
			log.Fatalf("Could not start the server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	<-quit

	log.Println("Server is stopping...")

	ctx, shutdown := context.WithTimeout(context.Background(), 5 * time.Second)
	defer shutdown()

	return s.echo.Shutdown(ctx)
}