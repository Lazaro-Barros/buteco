package main

import (
	"github.com/Lazaro-Barros/buteco/command/container"
	"log"
)

func main() {
	container.Init()
	r := container.GetRouter()
	container.GetDatabase()
	if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error to start server: %v", err)
    }
}
