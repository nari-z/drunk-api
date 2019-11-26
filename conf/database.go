package conf

// TODO: confから移動
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDBConnection(conf *Config) *gorm.DB {
	return getMysqlConn(conf)
}

func getMysqlConn(conf *Config) *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Database,
	)

	conn, err := gorm.Open("mysql", connectionString)
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

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
