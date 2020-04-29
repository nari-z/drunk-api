package conf

// TODO: confから移動
import (
	"fmt"

	// postgres driver.
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

// NewDBConnection return *gorm.DB.
func NewDBConnection(conf *Config) *gorm.DB {
	return getDBConn(conf)
}

// getDBConn return *gorm.DB.
func getDBConn(conf *Config) *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Database,
		conf.Database.Password,
	)

	conn, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(20)

	return conn
}
