package main

import (
	"backend/loaders"
)

func main() {
	loaders.DatabaseInit()
	loaders.InitRoutes()
}
