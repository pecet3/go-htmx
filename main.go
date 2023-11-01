package main

import (
	"github.com/pecet3/go+htmx/pkg/model"
	"github.com/pecet3/go+htmx/pkg/routes"
)

func main() {
	model.ConnectDb()

	routes.SetupAndRun()
}
