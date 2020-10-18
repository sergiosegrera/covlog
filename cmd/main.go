package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/kevinburke/twilio-go"
	"github.com/sergiosegrera/covlog/db/redisdb"
	"github.com/sergiosegrera/covlog/service"
	"go.uber.org/zap"
)

func main() {
	// Start logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	db, err := redisdb.New(conf)
	if err != nil {
		logger.Fatal("Error connecting to db")
	}

	tc, err := twilio.NewClient("", "", nil)
	if err != nil {
		logger.Fatal("Error creating twilio client")
	}

	covlogService := &service.CovlogService{
		DB:     db,
		tc:     tc,
		logger: logger,
	}

	go func() {
		logger.Info("Starting the http server", zap.String("port", "8080"))
		err := http.Serve(covlogService)
		if err != nil {
			logger.Error("Http server panic", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("exited", zap.String("sig", sig.String()))
}
