package main

import (
	"IDT-messaging/core/endpoints"
	"IDT-messaging/core/repository"
	"IDT-messaging/core/services"
	transport "IDT-messaging/core/transports/http"
	"IDT-messaging/core/transports/http/consts"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	httpPort := flag.String("port", consts.DefaultServicePort, "http port to listen on")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")

	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()

	ctx := context.Background()

	var usersRepo services.UsersRepo
	{
		usersRepo = repository.NewUsersRepo()
	}

	var service services.Service
	{
		service = services.NewService(logger, usersRepo)
	}

	errs := make(chan error)

	//listen for termination from the OS
	//e.g. user presses control+Z
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := endpoints.MakeEndpoints(service)

	//listen on endpoints
	//using transports specific handlers
	go func() {
		fmt.Printf("service listening on port [%v]", *httpPort)
		httpHandler := transport.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpPort, httpHandler)
	}()

	//force wait for termination
	level.Error(logger).Log("exit", <-errs)
}
