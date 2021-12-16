package database

import (
	redis "github.com/go-redis/redis/v8"
	"github.com/renjingneng/goapp/core/config"
	"strings"
)

var redisContainer map[string]interface{}

//GetEntityFromRedisContainer is
func GetEntityFromRedisContainer(database string, mode string) interface{} {
	if database == "" || mode == "" {
		return nil
	}
	dbname := database + mode
	if db, ok := redisContainer[dbname]; ok {
		return db
	}
	if ok := config.Get(dbname); ok == "" {
		return nil
	}
	var db interface{}
	if mode == "Single" {
		db = redis.NewClient(&redis.Options{
			Addr:     config.Get(dbname),
			Password: "",
			DB:       0,
		})
	} else if mode == "Cluster" {
		addrs := strings.Split(config.Get(dbname), ",")
		db = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: addrs,
		})
	}
	redisContainer[dbname] = db
	return db
}
func init() {
	if redisContainer == nil {
		redisContainer = make(map[string]interface{})
	}
}
