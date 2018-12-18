package main

import (
	"fmt"
	"os"
	// "github.com/radding/retroPieManager/internal/server"
	// "github.com/radding/retroPieManager/internal/update"
)

func main() {
	// update.IfNeeded()
	// server.StartServer()
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", name)
}
