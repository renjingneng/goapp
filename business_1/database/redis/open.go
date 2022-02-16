package redis

import (
	"github.com/renjingneng/goapp/core/database"
)

type Open struct {
	database.RedisBase
}

func NewOpen() *Open {
	return &Open{
		RedisBase: database.NewRedisBase("RedisOpen"),
	}
}
