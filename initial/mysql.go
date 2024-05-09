package initial

import (
	"fmt"
	"gorm.io/gorm"
	"my-douyin-fighting/glob"
)
import "gorm.io/driver/mysql"

func Mysql() {
	c := glob.Config.MysqlConfig
	username, password, host, port, dbname, maxopenconns, maxidelconns := c.Username, c.Password, c.Host, c.Password, c.DBName, c.MaxOpenConns, c.MaxIdleConns
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)
	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: nil,
	})
	glob.DB = db
	sql, _ := db.DB()
	sql.SetMaxOpenConns(maxopenconns)
	sql.SetMaxIdleConns(maxidelconns)
	if err != nil {
		return
	}
	if glob.AutoCreateDB {

	}

}
