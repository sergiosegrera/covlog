package redisdb

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sergiosegrera/covlog/config"
	"github.com/sergiosegrera/covlog/db"
	"github.com/sergiosegrera/covlog/models"
)

type RedisDB struct {
	client redis.Conn
}

func New(conf *config.Config) (db.DB, error) {
	client, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		return nil, err
	}

	return &RedisDB{
		client: client,
	}, err
}

func (r *RedisDB) SavePerson(ctx context.Context, p models.Person) error {
	// TODO: Being able to change default timeout.
	_, err := r.client.Do("SETEX", p.Phone, 1209600, p.Name)
	if err != nil {
		return err
	}

	_, err = r.client.Do("SADD", "phones", p.Phone)
	if err != nil {
		return err
	}

	return err
}

// TODO: Iterator and Count arguments, return iterator
func (r *RedisDB) GetPersons(ctx context.Context) ([]models.Person, error) {
	values, err := redis.Values(r.client.Do("SSCAN", "phones", 0, "COUNT", 5))
	if err != nil {
		return nil, err
	}

	iter, _ := redis.Int(values[0], nil)
	fmt.Println(iter)
	phones, _ := redis.Strings(values[1], nil)

	persons := []models.Person{}

	for _, phone := range phones {
		value, err := redis.String(r.client.Do("GET", phone))
		if err != nil {
			return nil, err
		}

		ttl, err := redis.Int(r.client.Do("TTL", phone))
		persons = append(persons, models.Person{
			Name:  value,
			Phone: phone,
			TTL:   ttl,
		})
		if err != nil {
			return nil, err
		}
	}

	return persons, err
}
