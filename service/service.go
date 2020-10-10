package service

import (
	"context"
	"errors"
	"github.com/kevinburke/twilio-go"
	"github.com/sergiosegrera/covlog/db"
	"github.com/sergiosegrera/covlog/models"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	CreatePerson(context.Context, models.Person) error
	// TODO: Maybe pagination? Probably pagination.
	GetPersons(context.Context) ([]models.Person, error)
	SendMessages(context.Context, string) error
}

type CovlogService struct {
	DB     db.DB
	tc     *twilio.Client
	logger *zap.Logger
}

func (s *CovlogService) CreatePerson(ctx context.Context, p models.Person) (err error) {
	defer func(begin time.Time) {
		s.logger.Info(
			"covlog",
			zap.String("method", "createperson"),
			zap.String("name", p.Name),
			zap.String("phone", p.Phone),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	err = s.DB.SavePerson(ctx, p)
	if err != nil {
		return ErrDatabase
	}

	// TODO: Send confirmation sms.
	_, err = s.tc.Messages.SendMessage("from", p.Phone, "Your number has been logged! It will automatically be deleted in 14 days.", nil)

	if err != nil {
		return ErrSendingMessage
	}

	return err
}

func (s *CovlogService) GetPersons(ctx context.Context) (persons []models.Person, err error) {
	defer func(begin time.Time) {
		s.logger.Info(
			"covlog",
			zap.String("method", "getpersons"),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	persons, err = s.DB.GetPersons(ctx)
	if err != nil {
		return nil, ErrDatabase
	}

	return persons, err
}

func (s *CovlogService) SendMessages(ctx context.Context, m string) (err error) {
	// TODO: logging.

	_, err = s.DB.GetPersons(ctx)
	if err != nil {
		return ErrDatabase
	}

	// TODO: Loop persons and send warning sms.

	return nil
}

var (
	ErrDatabase       = errors.New("Database error")
	ErrSendingMessage = errors.New("There was an error sending the confirmation message")
)
