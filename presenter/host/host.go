package host

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/nari-z/drunk-api/bundler"
	"github.com/nari-z/drunk-api/conf"
)

// NewHost is create new server.
func NewHost() {
	e := echo.New()

	var c *conf.Config = conf.NewConfig()
	var dbConn *gorm.DB = conf.NewDBConnection(c)
	var b *bundler.Bundler = bundler.NewBundler(dbConn)

	NewRouter(e, b.Handle)

	// static file.
	// TODO: ディレクトリの設定を全体と共有したい。
	e.Static("/LiquorImage", "LiquorImage")

	port := 1234
	err := e.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}
}
