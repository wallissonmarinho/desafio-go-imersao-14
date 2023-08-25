package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"

	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/endpoint"
	"github.com/wallissonmarinho/desafio-go-imersao-14/internal/service"
	transHttp "github.com/wallissonmarinho/desafio-go-imersao-14/internal/transport/http"
)

func main() {
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = kitlog.With(logger,
			"ts", kitlog.DefaultTimestampUTC,
			"caller", kitlog.DefaultCaller,
		)
	}

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes")
		if err != nil {
			os.Exit(-1)
		}
	}
	defer db.Close()

	level.Info(logger)
	defer level.Info(logger)

	var (
		context  context.Context
		services = service.NewServiceFactory(db, logger)
		endpoint = endpoint.MakeEndpoints(services, logger)
		err      = make(chan error)
	)

	go func() {
		server := &http.Server{
			Addr:    fmt.Sprintf(":%s", "8080"),
			Handler: transHttp.NewService(context, &endpoint, &logger),
		}
		err <- server.ListenAndServe()
	}()

	fatal := level.Error(logger).Log("exit", <-err)
	if fatal != nil {
		logrus.Error(fatal)
	}

}
