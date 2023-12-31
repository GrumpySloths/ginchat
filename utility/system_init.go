package utility

import (
	"fmt"
	"ginchat/models"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB       *gorm.DB
	RedisCli *redis.Client
)

func Config_init() {
	viper.SetConfigName("system_init")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println(viper.GetString("mysql.dsn"))
}

func Redis_init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.DB"),
	})
	pong, err := RedisCli.Ping().Result()
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println(pong)
	}

}

func DataBase_init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,             // Don't include params in the SQL log
			Colorful:                  true,             // Disable color
		},
	)
	dsn := viper.GetString("mysql.dsn")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	DB.AutoMigrate(&models.UserBasic{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.GroupBasic{})
	DB.AutoMigrate(&models.User_Group{})
}
