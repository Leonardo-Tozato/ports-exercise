package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"ports-exercise/m/internal/handlers/portshdl"
	"ports-exercise/m/internal/port"
	"ports-exercise/m/internal/repositories/portsrepo"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}
	filePath := os.Args[1]

	// Set up signal handling
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	// Set up dependencies
	portsRepository := portsrepo.NewMemDB()
	portsService := port.NewService(portsRepository)
	portsHandler := portshdl.NewJsonHandler(portsService)

	// Process input json
	result, err := portsHandler.HandleUpsertStream(filePath)
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	<-shutdown
	fmt.Println("DONE")
}
