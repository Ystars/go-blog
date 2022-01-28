package main

import (
	"goblog/config"
	"goblog/routes"
	"os"
)

func main() {

	config.Init()

	r := routes.Load()

	_ = r.Run(os.Getenv("HTTP_PORT"))
}
