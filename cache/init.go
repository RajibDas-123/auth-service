package cache

import (
	"os"
	"runtime"

	"github.com/RajibDas-123/ms-grpc-auth/auth/logging"

	"github.com/go-redis/redis"
)

var JobConnection *redis.Client
var PubSub *redis.Client

func getConnection(db int) *redis.Client {
	var redisConnection = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("CACHE_SERVER_ADDR"),
		Password:     os.Getenv("CACHE_PASSWORD"),
		DB:           db,
		MaxRetries:   5,
		PoolSize:     10 * runtime.NumCPU(),
		MinIdleConns: 5,
	})
	if _, err := redisConnection.Ping().Result(); err != nil {
		logging.CacheLogger.Fatal("Failed to connect to the database. ", err)
	}
	return redisConnection
}

func Initialize() {
	JobConnection = getConnection(0)
	PubSub = getConnection(1)

	logging.CacheLogger.Info("Cache connections established")
}
