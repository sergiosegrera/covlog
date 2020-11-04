package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/kevinburke/twilio-go"
	"github.com/sergiosegrera/covlog/config"
	"github.com/sergiosegrera/covlog/db/redisdb"
	"github.com/sergiosegrera/covlog/service"
	"github.com/sergiosegrera/covlog/transports/http"
	"go.uber.org/zap"
)

func main() {
	// Load config
	conf := config.New()
	fmt.Println(conf)

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

	// TODO: Check to see if connection fails
	tc := twilio.NewClient(conf.TwilioId, conf.TwilioToken, nil)

	covlogService := &service.CovlogService{
		DB:        db,
		TC:        tc,
		Logger:    logger,
		FromPhone: conf.TwilioPhone,
	}

	go func() {
		logger.Info("Starting the http server", zap.String("port", conf.HttpPort))
		err := http.Serve(covlogService, conf)
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
