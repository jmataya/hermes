package main

import (
	"os"

	"github.com/jmataya/hermes/srv"
	"github.com/jmataya/hermes/utils"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hermes"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start the server",
			Action: func(c *cli.Context) error {
				srv.Run()
				return nil
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "Run database migrations",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "database",
					Usage: "Database name to migrate",
				},
				cli.StringFlag{
					Name:  "host",
					Value: "localhost",
					Usage: "Hostname containing the database instance to migrate",
				},
				cli.StringFlag{
					Name:  "user",
					Value: "postgres",
					Usage: "User to use when connecting to the database",
				},
				cli.StringFlag{
					Name:  "password",
					Usage: "Password to use when connecting to the database",
				},
				cli.BoolFlag{
					Name:  "ssl",
					Usage: "Enable SSL with the database connection",
				},
				cli.StringFlag{
					Name:  "source",
					Value: "sql",
					Usage: "Location of the SQL migration files",
				},
			},
			Action: func(c *cli.Context) error {
				database := c.String("database")
				if database == "" {
					return cli.NewExitError("database parameter must be specified", 1)
				}

				err := utils.MigratePG(
					c.String("source"),
					c.String("host"),
					c.String("user"),
					c.String("password"),
					c.String("database"),
					c.Bool("ssl"))

				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
