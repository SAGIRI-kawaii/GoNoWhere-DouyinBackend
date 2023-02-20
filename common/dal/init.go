package dal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mini-douyin/common/dal/config"
	"time"
)

var DB *gorm.DB

func InitDB() {

	// 手动编写组成数据库连接串
	config.InitConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.charset"),
	)
	fmt.Println(dsn)
	//dsn2 := "where:where@tcp(8.142.30.177:3306)/gonowhere?charset=utf8mb4&parseTime=True"
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
		// util.LogrusObj.Error(err)
	}

}
func Database(dsn string) error {
	var ormlogger logger.Interface
	if gin.Mode() == "debug" {
		ormlogger = logger.Default.LogMode(logger.Info)
	} else {
		ormlogger = logger.Default
	}
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			// PrepareStmt: true,
			// SkipDefaultTransaction: true,
			Logger: ormlogger,
		},
	)
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db

	// migration()
	return err

}
