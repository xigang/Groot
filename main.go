package main

import (
	"os"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xigang/groot/cmd"
	"github.com/xigang/groot/pkg/kubernetes"
)

func main() {
	app := cli.NewApp()

	app.Name = "groot"
	app.Usage = "groot is the command line tool for kubeflow"
	app.Version = "0.1.0"
	app.Author = "wangxigang2014@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "/etc/kubernetes/kubeconfig",
			Usage: "path to a kube config. only required if out-of-cluster",
		},
	}

	app.Before = func(c *cli.Context) error {
		kubeconfig := c.GlobalString("config")

		clientset, err := kubernetes.CreateKubeClient(kubeconfig)
		if err != nil {
			logrus.Errorf("failed to create kube client: %v", err)
			return err
		}

		kubernetes.KubeClient = clientset

		return nil
	}

	app.Commands = []cli.Command{
		cmd.GrootVersion,
		cmd.SubmitAction,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
