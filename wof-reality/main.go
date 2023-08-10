package main

import (
	"fmt"
	"os"

	"wof/wof-reality/service"

	"github.com/cloudfoundry-community/go-cfenv"
)

func main() {
	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Printf("CF environment not detected. APP WILL RUN WITH FAKE REPOSITORY!\n")
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer(appEnv)
	server.Run(":" + port)
}
