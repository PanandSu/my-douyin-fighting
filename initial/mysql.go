package initial

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
	"os"
	"time"
)

func Mysql() {
	c := gb.Cfg.MysqlConfig
	username, password, host, port, dbname :=
		c.Username, c.Password, c.Host, c.Port, c.DBName

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 0,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		})
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("connect mysql error:" + err.Error())
	}
	gb.DB = db
	sql, _ := db.DB()
	sql.SetMaxOpenConns(gb.Cfg.MysqlConfig.MaxOpenConns)
	sql.SetMaxIdleConns(gb.Cfg.MysqlConfig.MaxIdleConns)
	if gb.AutoCreateDB {
		gb.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
		gb.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Video{})
		gb.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Like{})
		gb.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Follow{})
		gb.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Comment{})
	}
}
