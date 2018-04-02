package main

import (
	"fmt"
	"os"

	"github.com/CardInfoLink/log"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"

	"github.com/wonsikin/beehive/src"
	"github.com/wonsikin/beehive/src/config"
	"github.com/wonsikin/beehive/src/db"
	"github.com/wonsikin/beehive/src/server"
)

// constant of app
const (
	AppVersion  = "v0.1.0"
	AppName     = "beehive"
	DefaultPort = 13000
)

var (
	defaultWorkerCfgFilePath = fmt.Sprintf("./%s", config.WorkerCfgFileName)
	defaultQueenCfgFilePath  = fmt.Sprintf("./%s", config.QueenCfgFileName)
)

var validRoles = []string{config.QueenRole, config.WorkerRole}

func main() {
	app := cli.NewApp()
	app.Version = AppVersion
	app.Name = AppName
	app.HelpName = AppName
	app.Usage = "eat log messages and display it"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "role, r",
			Value: "queen",
			Usage: fmt.Sprintf("running as '%s' or '%s'", config.QueenRole, config.WorkerRole),
		},
		cli.IntFlag{
			Name:  "port, p",
			Value: DefaultPort,
			Usage: fmt.Sprintf("listen port of %s", AppName),
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
			Usage: fmt.Sprintf("load configuration from `FILE`, default value is %s when running as the queen role and is %s when running as worker role", defaultQueenCfgFilePath, defaultWorkerCfgFilePath),
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "create a configuration file",
			Action: func(c *cli.Context) error {
				prompt := promptui.Select{
					Label: "Generate configuration for which role?",
					Items: validRoles,
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
		// role is required and should handle first
		role := c.String("role")
		switch role {
		case config.QueenRole:
			return runQueenServer(c)
		case config.WorkerRole:
			return runWorker(c)
		default:
			return fmt.Errorf("invalid role")
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
	}
}

func runQueenServer(c *cli.Context) error {
	// parse config file
	path := c.String("config")
	if path == "" {
		path = defaultQueenCfgFilePath
	}
	cfg, err := config.ParseQueenCfg(path)
	if err != nil {
		return err
	}

	// init mongo connection
	err = db.Connect(cfg.DB)
	if err != nil {
		log.Errorf("error caught when connecting to database(%s), error is %s", cfg.DB.Type, err)
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

func runWorker(c *cli.Context) error {
	// read configuartion file
	path := c.String("config")
	if path == "" {
		path = defaultWorkerCfgFilePath
	}
	cfg, err := config.ParseWorkerCfg(path)
	if err != nil {
		log.Errorf("errot caught when parsing configuration file: %s", err)
		return err
	}

	// start worker
	worker, err := src.NewWorker(cfg)
	if err != nil {
		log.Errorf("errot caught when creating worker: %s", err)
		return err
	}
	worker.Run()

	return nil
}
