package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var ideas Ideas
	ideas.Load()
	app := &cli.App{
		Name: "idea",
		Usage: "A lighweight CLI tool for keeping your ideas. " +
			"This is a cloned version of idea CLI tool written in golang. " +
			"https://github.com/rmsubekti/idea",
		UsageText: "idea [command] <idea|state|id>",
		Action: func(c *cli.Context) error {
			ideas.ListByState("OPEN")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Create new .ideas.json file in the current directory. Default: ~/.ideas.json",
				Action: func(c *cli.Context) error {
					return Init()
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new idea. Example: `idea create CLI app`",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return errors.New("Need one or more arguments." +
							"\n\tTry `idea create \"string argument\"` or `idea create some idea text`" +
							"\n\tor type \"idea help create\" for help")
					}
					idea := strings.Join(c.Args().Slice(), " ")
					return ideas.Create(idea)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete an idea. Example: `idea delete 1`",
				Action: func(c *cli.Context) error {
					id, err := strconv.Atoi(c.Args().First())
					err = ideas.Remove(id)
					return err
				},
			},
			{
				Name:    "solve",
				Aliases: []string{"s"},
				Usage:   "Solve an idea. Example: `idea solve 1`",
				Action: func(c *cli.Context) error {
					id, err := strconv.Atoi(c.Args().First())
					ideas.Solve(id)
					return err
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all ideas. Example `idea list solved`.(default: open)",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						ideas.ListByState(c.Args().First())
					} else {
						ideas.List()
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
