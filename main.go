package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/myselfBZ/user-service/env"
	"github.com/myselfBZ/user-service/service"
	"github.com/myselfBZ/user-service/store"
)

type config struct {
	port        int
	servicePort int
	db          store.DBConfig
}

func (c *config) Load() {
	c.port = env.GetInt("PORT",  8080)
	c.servicePort = env.GetInt("SERVICE_PORT", 6767)
	c.db.Addr = env.MustGetString("DB")
	c.db.MaxIdleConns = env.GetInt("DB_MAX_IDLE_CONNS", 30)
	c.db.MaxOpenConns = env.GetInt("DB_MAX_OPEN", 30)
	c.db.MaxIdleTime = env.GetString("DB_MAX_IDLE_TIME", "15m")
}
func main() {
	cfg := &config{}
	cfg.Load()
	db, err := store.NewDBConn(cfg.db)
	if err != nil {
		log.Fatalf("could not establish a database connection: %v", err)
	}
	st := store.NewPostgreStore(db)
	service := service.New(st)
	serviceAddr := fmt.Sprintf(":%d", cfg.servicePort)
	log.Println("let's fuck")
	log.Fatalf("service failure: %v", service.Run(serviceAddr))
}
