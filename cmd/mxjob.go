package cmd

import "github.com/urfave/cli"

var MXJobSubCommand = cli.Command{
	Name:    "mxjob",
	Aliases: []string{"mx"},
	Usage:   "submit a MXJob as training job.",
	Action: func(c *cli.Context) error {
		return nil
	},
}
