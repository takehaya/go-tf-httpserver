package main

import (
	"log"
	"os"

	"github.com/takehaya/go-tf-httpserver/internal"
	"github.com/takehaya/go-tf-httpserver/pkg/version"
)

func main() {
	app := internal.NewApp(version.Version)
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%+v", err)
	}
}
