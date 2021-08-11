package main

import (
	"fmt"
	"os"
	"time"

	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/authfun/gauthfun/config"
	"github.com/authfun/gauthfun/router"
)

// Graceful shutdown
// https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
// https://github.com/gin-gonic/examples/tree/master/graceful-shutdown
func main() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConfig.Service.Port),
		Handler: router.NewRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
