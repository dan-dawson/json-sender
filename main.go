package main

import (
	"fmt"
	"github.com/dan-daws/json-sender/api"
)

func main() {
	err := api.StartServer()
	if err != nil {
		fmt.Println("server failed to start correctly. Exiting...")
	}
}
