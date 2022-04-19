package data

import (
	kratos_log "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"kratos-demo/internal/conf"
	"log"
	"os"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewGreeterRepo, NewUserRepo)

type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger kratos_log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		kratos_log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "m_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}
	InitDB(db)
	return db
}
func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}, &Order{}); err != nil {
		panic(err)
	}
}
