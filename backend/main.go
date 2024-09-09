package main

import (
	"github.com/Lazaro-Barros/buteco/command/container"
)

func main() {
	container.Init()
	r := container.GetRouter()
	r.Run(":8080")

}
