package api

import (
	"fmt"
	"net/http"
	"os"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("you messed up fool")
	}

	w.Write(file)
}

func StartServer() error {
	http.HandleFunc("/json", jsonHandler)
	err := http.ListenAndServe(":6000", nil)

	return err
}
