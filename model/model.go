package model

import (
	"fmt"
	cache "gitea.com/xorm/xorm-redis-cache"
	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"hextechdocs-be/configuration"
	"hextechdocs-be/logger/integrations"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var (
	x      *xorm.Engine
	tables []interface{}
)

func initTables() {
	tables = append(tables,
		new(Author),
		new(Category),
		new(Document),
		new(Marker),
		new(Subcategory),
	)
}

func initCaching(config *configuration.Configuration) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisUrl)
			if err != nil {
				return nil, err
			}
			if config.ShouldUseRedisAuth {
				if _, err := c.Do("AUTH", config.RedisPassword); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			if config.ShouldUseRedisDatabase {
				if _, err := c.Do("SELECT", config.RedisDatabase); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
	x.SetDefaultCacher(cache.MakeRedisCacher(pool, time.Duration(10)*time.Minute, x.Logger()))
}

func IsDatabaseAlive() (bool, error) {
	ping := x.Ping()
	return x.Ping() == nil, ping
}

func Init(config *configuration.Configuration) error {
	initTables()

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName)

	var err error
	x, err = xorm.NewEngine("postgres", connectionString)
	if err != nil {
		return err
	}

	x.SetLogger(log.NewLoggerAdapter(&integrations.XormLogger{}))

	err = x.Sync2(tables...)
	if err != nil {
		return err
	}

	if config.ShouldUseRedis {
		initCaching(config)
	}

	return nil
}
