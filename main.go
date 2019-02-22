package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xigang/groot/cmd"
	tfjobclientset "github.com/xigang/groot/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
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
		kcfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return err
		}

		cs, err := tfjobclientset.NewForConfig(kcfg)
		if err != nil {
			return err
		}

		TFjob := cs.KubeflowV1beta2().TFJobs("default")
		result, err := TFjob.Get("tfjob-20190222", v1.GetOptions{})
		if err != nil {
			return err
		}

		fmt.Println(result)

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
