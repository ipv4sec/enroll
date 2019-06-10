package redis

import (
	"enroll/config"
	"enroll/logger"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"time"
)

var Clinet *redis.Client

func Init() {
	Clinet = redis.NewClient(config.Conf.Redis)

	_, err := Clinet.Ping().Result()
	if err != nil {
		logger.Error("Connect Redis Fail:", err.Error())
		os.Exit(0)
	}

	logger.Info("Connected To Redis:", config.Conf.Redis.Addr)
}

func SaveUid(token string, uid int64) error {
	return Clinet.Set("enroll:token:" + token, uid, time.Hour * 24).Err()
}

func FindUidByToken(token string) (int64, error) {
	result, err := Clinet.Get("enroll:token:" + token).Result()
	if err != nil {
		return 0, nil
	}
	return strconv.ParseInt(result, 10, 64)
}