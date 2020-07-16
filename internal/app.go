package internal

import (
	"log"

	"github.com/takehaya/go-tf-httpserver/pkg/utils"
	"github.com/urfave/cli"
)

func NewApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "go-tf-httpserver"
	app.Version = version

	app.Usage = "Services for using terraform http backend"

	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "./scripts/config.yaml",
			Usage: "setup config",
		},
	}
	app.Action = run
	return app
}

func run(ctx *cli.Context) error {
	configpath := ctx.String("config")

	if !utils.FileExists(configpath) {
		log.Fatalf("config file not found: %s", configpath)
	}
	err := Main(&configpath)
	if err != nil {
		return err
	}

	return nil
}

func Main(configpath *string) error {

	return nil
}
