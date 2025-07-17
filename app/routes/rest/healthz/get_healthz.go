package healthz

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type getHealthzAPIResponse struct {
	Status dependenciesStatus `json:"status"`
}

type statusEnums string

var (
	Ok    statusEnums = "OK"
	NotOk statusEnums = "NOT_OK"
)

type dependenciesStatus struct {
	Postgres statusEnums `json:"postgres"`
	Redis    statusEnums `json:"redis"`
}

func getHealthz(
	router fiber.Router,
	gormDB *gorm.DB,
	redisClient *redis.Client,
) {
	url := fmt.Sprintf("%s/check", prefix)
	router.Get(url, func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		resp := &dependenciesStatus{
			Postgres: Ok,
			Redis:    Ok,
		}

		sqlDB, err := gormDB.DB()
		if err != nil || sqlDB.Ping() != nil {
			log.Error().Err(err).Msg("Unable connect to postgres")
			resp.Postgres = "NOT_OK"
		}

		err = redisClient.Ping(ctx).Err()
		if err != nil {
			log.Error().Err(err).Msg("Unable connect to redis client")
			resp.Redis = "NOT_OK"
		}

		return c.JSON(&getHealthzAPIResponse{
			Status: *resp,
		})
	})
}
