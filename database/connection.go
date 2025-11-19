package database

import (
	"log"

	"github.com/abdallahelassal/Store/config"
	"github.com/abdallahelassal/Store/internal/modules/user/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
) 


type Connection struct{
	Cfg 	*config.Config
	DB		*gorm.DB
	Log		*log.Logger
}

func NewConnection(cfg *config.Config,  log *log.Logger)*Connection{
	conn := &Connection{
		Cfg: cfg,
		Log: log,
	}
	conn.Connection()
	return conn
}

func (c Connection) Connection(){
	dsn := "host="+c.Cfg.DB_HOST+
	"user="+c.Cfg.DB_USER+
	"password="+c.Cfg.DB_PASSWORD+
	"dbname="+c.Cfg.DB_NAME+
	"port="+c.Cfg.PORT+
	"sslmode=disable"
		

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
		panic("Failed to connect to database!")
	}
	c.DB = db
	c.DB.AutoMigrate(&domain.User{})
}
