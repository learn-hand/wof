package main

import (
	"os"

	"github.com/cloudnativego/wof-commandprocessing/service"
)

func main() {
	// appEnv, err := cfenv.Current()
	// if err != nil {
	// 	fmt.Printf("Failed to get a CF environment.\n")
	// 	os.Exit(1)
	// }

	dispatchers := make(map[string]service.QueueDispatcher)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer(nil, dispatchers)
	server.Run(":" + port)
}
