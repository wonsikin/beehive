package main

import (
	"fmt"
	"os"

	"github.com/CardInfoLink/log"
	"github.com/urfave/cli"

	"github.com/wonsikin/beehive/worker/src"
	"github.com/wonsikin/beehive/worker/src/cfg"
)

// constant of app
const (
	AppVersion = "v0.0.1"
	AppName    = "bh-worker"
)

func main() {
	app := cli.NewApp()
	app.Version = AppVersion
	app.Name = AppName
	app.HelpName = AppName
	app.Usage = "eat log files and squeeze out messages"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: fmt.Sprintf("./%s", cfg.CfgFileName),
			Usage: "load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "create a configuration file",
			Action: func(c *cli.Context) error {
				err := cfg.Init()
				if err != nil {
					log.Errorf("errot caught when initing configuration file: %s", err)
					return err
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		cPath := c.String("config")
		// read configuartion file
		fmt.Printf("config path is %s\n", cPath)
		config, err := cfg.Parse(cPath)
		if err != nil {
			log.Errorf("errot caught when parsing configuration file: %s", err)
			return err
		}

		// start worker
		worker, err := src.NewWorker(config)
		if err != nil {
			log.Errorf("errot caught when creating worker: %s", err)
			return err
		}
		worker.Run()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
	}
}
