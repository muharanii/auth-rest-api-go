package component

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"log"
	"github.com/muharanii/auth-rest-api-go/domain"
	"time"
)

func GetCacheConnection() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		log.Fatalf("error when connect cache %s", err.Error())
	}
	return cache
}
