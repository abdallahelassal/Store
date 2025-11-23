package database

import (


	"github.com/abdallahelassal/Store/config"
	"github.com/abdallahelassal/Store/internal/modules/user/domain"
	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
) 


type Connection struct{
	Cfg 	*config.Config
	DB		*gorm.DB
	Log		*zap.Logger
}

func NewConnection(cfg *config.Config,  log *zap.Logger)*Connection{
	conn := &Connection{
		Cfg: cfg,
		Log: log,
	}
	conn.Connection()
	return conn
}

func (c *Connection) Connection(){
	dsn := "host="+c.Cfg.DatabaseConfig.Host+
		" user="+c.Cfg.DatabaseConfig.User+
		" password="+c.Cfg.DatabaseConfig.Password+
		" dbname="+c.Cfg.DatabaseConfig.Name+
		" port="+c.Cfg.DatabaseConfig.Port+
		" sslmode="+c.Cfg.DatabaseConfig.SSLMode
		
	c.Log.Info("Connecting to database",
		zap.String("host", c.Cfg.DatabaseConfig.Host),
		zap.String("user", c.Cfg.DatabaseConfig.User),
		zap.String("dbname", c.Cfg.DatabaseConfig.Name),
		zap.String("port", c.Cfg.DatabaseConfig.Port),
		zap.String("sslmode", c.Cfg.DatabaseConfig.SSLMode),
	)


	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
		panic("Failed to connect to database!")
	}
	c.DB = db
	if err := c.DB.AutoMigrate(&domain.User{}); err != nil {
		c.Log.Info("migration err:")
		return
	}
}

func (c *Connection) Close(){
	sqlDB, err := c.DB.DB()
	if err == nil {
		sqlDB.Close()
	}
}