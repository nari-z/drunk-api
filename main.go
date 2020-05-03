package main

import (
	"log"

	"github.com/nari-z/drunk-api/presenter/host"
)

func main() {
	server, err := host.NewHost()
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Shutdown()

	err = server.Serve();
	if err != nil {
		log.Fatalln(err)
	}
}
