package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	image_service_server "github.com/ritik-patel05/image-player/internal/app/image-service"
)

const (
	port = ":9876"
)

func main() {
	var srv *http.Server

	idleConnClosed := make(chan struct{})
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-done

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		image_service_server.Shutdown(ctx, srv)
		defer cancel()
		close(idleConnClosed)
	}()

	router := image_service_server.NewServer()
	srv = &http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Server started on port", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("Server Closed")
	}
	<-idleConnClosed
}
