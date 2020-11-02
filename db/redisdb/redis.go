package redisdb

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sergiosegrera/covlog/db"
	"github.com/sergiosegrera/covlog/config"
	"github.com/sergiosegrera/covlog/models"
)

type RedisDB struct {
	client *redis.Client
}

func New(conf *config.Config) (db.DB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
		DB:       0,
	})

	// TODO: Add timeout.
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisDB{
		client: client,
	}, err
}

func (r *RedisDB) SavePerson(ctx context.Context, p models.Person) error {
	// TODO: Being able to change default timeout.
	err := r.client.Set(ctx, "phone:"+p.Phone, p.Name, 14*24*time.Hour).Err()

	return err
}
