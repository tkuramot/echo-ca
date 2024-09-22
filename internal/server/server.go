package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/tkuramot/echo-practice/internal/config"
	"github/tkuramot/echo-practice/internal/presentation/settings"
	"github/tkuramot/echo-practice/internal/server/route"
)

func Run(ctx context.Context, conf *config.Config) {
	e := settings.NewEcho()
	route.InitRoute(e)

	address := conf.Server.Address + ":" + conf.Server.Port
	log.Println("Server is running at " + address)
	srv := &http.Server{
		Addr:              address,
		Handler:           e,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
