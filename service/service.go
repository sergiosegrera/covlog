package service

import (
	"context"
	"errors"
	"time"

	"github.com/kevinburke/twilio-go"
	"github.com/sergiosegrera/covlog/db"
	"github.com/sergiosegrera/covlog/models"
	"go.uber.org/zap"
)

type Service interface {
	CreatePerson(context.Context, models.Person) error
	// TODO: Maybe pagination? Probably pagination.
	GetPersons(context.Context) ([]models.Person, error)
	SendMessages(context.Context, string) error
}

type CovlogService struct {
	DB        db.DB
	TC        *twilio.Client
	Logger    *zap.Logger
	FromPhone string
}

func (s *CovlogService) CreatePerson(ctx context.Context, p models.Person) (err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
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

	return err
}

func (s *CovlogService) GetPersons(ctx context.Context) (persons []models.Person, err error) {
	defer func(begin time.Time) {
		s.Logger.Info(
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
	var persons []models.Person
	// TODO: Log total sent count
	defer func(begin time.Time) {
		s.Logger.Info(
			"covlog",
			zap.String("method", "sendmessages"),
			zap.Int("sent", len(persons)),
			zap.NamedError("err", err),
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())

	// TODO: This is not scalable possibly use chanels?
	persons, err = s.DB.GetPersons(ctx)
	if err != nil {
		return ErrDatabase
	}

	// TODO: Catch different errors
	for _, person := range persons {
		s.TC.Messages.SendMessage(s.FromPhone, person.Phone, m, nil)
	}

	return nil
}

var (
	ErrDatabase       = errors.New("Database error")
	ErrSendingMessage = errors.New("There was an error sending the confirmation message")
)
