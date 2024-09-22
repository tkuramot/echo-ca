package main

import (
	"context"

	"github/tkuramot/echo-practice/internal/config"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db"
	"github/tkuramot/echo-practice/internal/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	db.NewMainDB(conf.DB)
	db.NewReadDB(conf.ReadDB)

	server.Run(ctx, conf)
}
