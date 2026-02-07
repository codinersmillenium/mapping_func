package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mapping_func/config"
	"mapping_func/core/repository"
	"mapping_func/core/service"
	"mapping_func/pkg/handler"
	route "mapping_func/pkg/http"
	"mapping_func/pkg/utils"
)

func main() {
	cfg := config.Load()

	db, err := config.NewDatabase(cfg.DBDriver, cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	cities, err := utils.LoadCities(cfg.CityPath)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cities)
	userHandler := handler.NewUserHandler(userService)

	// set gin mode
	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	// set trusted proxies
	if len(cfg.TrustedProxies) > 0 && !(len(cfg.TrustedProxies) == 1 && cfg.TrustedProxies[0] == "") {
		r.SetTrustedProxies(cfg.TrustedProxies)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     utils.CorsOrigins(cfg),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.Register(r, userHandler)
	r.Run(":" + cfg.AppPort)
}
