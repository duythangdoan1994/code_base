package main

import (
	"github.com/luciandd/core/api"
)

func main() {
	router := api.SetupRouter()
	router.Run(":8080")
}