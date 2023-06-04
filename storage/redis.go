package storage

import (
	"horus-watcher/configs"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func GetServicesNames() []string {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	v, err := Redis.HGetAll("defense-services").Result()
	if err != nil {
		panic(err)
	}
	var servicesNames []string
	for k := range v {
		servicesNames = append(servicesNames, k)
	}
	return servicesNames
}

func GetSubscribedUrls() []string {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	v, err := Redis.HGetAll("subscribe").Result()
	if err != nil {
		panic(err)
	}
	var subscribedUrls []string
	for k := range v {
		subscribedUrls = append(subscribedUrls, k)
	}
	return subscribedUrls
}

func GetFlaggedServices() []string {
	conf := configs.GetRedisConfig()
	Redis = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + conf.Port,
		// Password: conf.Password,
	})
	v, err := Redis.HGetAll("flagged-services").Result()
	if err != nil {
		panic(err)
	}
	var flaggedServices []string
	for k := range v {
		flaggedServices = append(flaggedServices, k)
	}
	return flaggedServices
}