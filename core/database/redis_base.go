package database

import (
	"context"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/renjingneng/goapp/core/config"
)

type RedisBase interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
}
type singleRedisBase struct {
	Dbname string
	Mode   string
	ctx    context.Context
	Dbptr  *redis.Client
}
type clusterRedisBase struct {
	Dbname string
	Mode   string
	ctx    context.Context
	Dbptr  *redis.ClusterClient
}

func NewRedisBase(Dbname string) RedisBase {
	redisType := config.Get(Dbname + "Type")
	dbptr := GetEntityFromRedisContainer(Dbname, redisType)
	if strings.ToUpper(redisType) == "SINGLE" {
		db := &singleRedisBase{
			Dbname: Dbname,
			Mode:   "Single",
			ctx:    context.Background(),
		}
		db.Dbptr, _ = dbptr.(*redis.Client)
		return db
	} else if strings.ToUpper(redisType) == "CLUSTER" {
		db := &clusterRedisBase{
			Dbname: Dbname,
			Mode:   "Cluster",
			ctx:    context.Background(),
		}
		db.Dbptr, _ = dbptr.(*redis.ClusterClient)
		return db
	} else {
		return nil
	}
}

func (b *singleRedisBase) Get(key string) (string, error) {
	val, err := b.Dbptr.Get(b.ctx, key).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}
func (b *singleRedisBase) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	val, err := b.Dbptr.Set(b.ctx, key, value, expiration).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}

func (b *clusterRedisBase) Get(key string) (string, error) {
	val, err := b.Dbptr.Get(b.ctx, key).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}

func (b *clusterRedisBase) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	val, err := b.Dbptr.Set(b.ctx, key, value, expiration).Result()
	if err != nil {
		return val, err
	} else {
		return val, nil
	}
}
