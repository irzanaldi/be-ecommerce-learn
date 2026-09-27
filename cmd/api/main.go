package main

import (
	"github.com/irzan-aldi/be-ecommerce/internal/config"
	"github.com/irzan-aldi/be-ecommerce/internal/router"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Failled Connect DB: ", err)
	}

	r := router.Setup(db, cfg)

	log.Printf("Server Running in Port : ", cfg.Port)
	r.Run(":" + cfg.Port)
}
