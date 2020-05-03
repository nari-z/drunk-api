package host

import (
	"github.com/jinzhu/gorm"
	"github.com/go-openapi/loads"

	"github.com/nari-z/drunk-api/bundler"
	"github.com/nari-z/drunk-api/conf"
	"github.com/nari-z/drunk-api/generate/restapi"
	"github.com/nari-z/drunk-api/generate/restapi/operations"
)

// NewHost is create new server.
func NewHost() (*restapi.Server, error) {
	// get swagger config
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewDrunkAPI(swaggerSpec)

	// get config
	var c *conf.Config = conf.NewConfig()
	var dbConn *gorm.DB = conf.NewDBConnection(c)
	var b *bundler.Bundler = bundler.NewBundler(dbConn)
	NewRouter(api, b.Handle)

	server := restapi.NewServer(api)
	server.ConfigureAPI()
	server.Port = 1234

	return server, nil



}
