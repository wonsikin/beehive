package main

import (
	"fmt"
	"os"

	"github.com/CardInfoLink/log"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"

	"github.com/wonsikin/beehive/queen/src/cfg"
	"github.com/wonsikin/beehive/queen/src/db"
	"github.com/wonsikin/beehive/queen/src/server"
	"github.com/wonsikin/beehive/src/config"
)

// constant of app
const (
	AppVersion  = "v0.1.0"
	AppName     = "beehive"
	DefaultPort = 13000
)

func main() {
	app := cli.NewApp()
	app.Version = AppVersion
	app.Name = AppName
	app.HelpName = AppName
	app.Usage = "eat messages and display it"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "role, r",
			Value: "queen",
			Usage: "running as 'queen bee' or 'worker bee'",
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: DefaultPort,
			Usage: fmt.Sprintf("listen port of %s", AppName),
		},
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
				prompt := promptui.Select{
					Label: "Generate configuration for which role?",
					Items: []string{"queen", "worker"},
				}

				_, choice, err := prompt.Run()
				if err != nil {
					return err
				}

				log.Debugf("selected item is %s", choice)
				return config.Init(choice)
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		// parse config file
		configFile := c.String("config")
		config, err := cfg.Parse(configFile)
		if err != nil {
			return err
		}
		log.Debugf("config is %#v", config)

		// init mongo connection
		err = db.Connect(config.DB)
		if err != nil {
			log.Errorf("error caught when connecting to database(%s), error is %s", config.DB.Type, err)
			return err
		}

		// TODO compute server address
		// compute whether the default port is occupied
		port := c.Int("port")
		address := fmt.Sprintf(":%d", port)
		srv := server.NewServer(address)

		log.Infof("%s is served at %s", AppName, address)
		log.Error(srv.ListenAndServe())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
	}
}