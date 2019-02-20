package cmd

import (
	"github.com/urfave/cli"
)

var SubmitAction = cli.Command{
	Name:  "submit",
	Usage: "submit a job",
	Subcommands: []cli.Command{
		TFJobSubCommand,
		MXJobSubCommand,
	},
}
