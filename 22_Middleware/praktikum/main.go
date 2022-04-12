package main

import (
	"praktikum/config"
	"praktikum/routes"
)

func main() {
	config.Init()
	e := routes.New()
	e.Start(":8002")
}
