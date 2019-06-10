package config

import "github.com/go-redis/redis"

type Config struct {
	Application struct{
		Address string
	}

	Mysql struct {
		Host string
		User string
		Pass string
		Port int
		Db string
		// user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	}

	Redis *redis.Options

	Csv struct{
		Example string
		Uploaded string
	}
}



