package conf

import (
	"os"
)

type Config struct {
	// TODO: network setting.
	// Server struct {
	// 	Port int
	// 	Mock bool
	// }
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

func NewConfig() *Config {
	var c *Config = &Config{};

	// データベース設定を取得
	c.Database.Host = os.Getenv("DRUNK_DATABASE_HOST");
	c.Database.Port = os.Getenv("DRUNK_DATABASE_PORT");
	c.Database.User = os.Getenv("DRUNK_DATABASE_USER");
	c.Database.Password = os.Getenv("DRUNK_DATABASE_PASSWORD");
	c.Database.Database = os.Getenv("DRUNK_DATABASE_NAME");

	return c;
}


