package main

import (
	"belajar-go-echo/route"
)

func main() {
	app := route.New()

	app.Start(":8080")
}
