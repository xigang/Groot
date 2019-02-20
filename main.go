package main

import (
	"os"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/xigang/groot/cmd"
	"github.com/xigang/groot/pkg/kubernetes"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		cli.StringFlag{
			Name:  "namespace",
			Value: "default",
			Usage: "the namespace of the job",
		},
	}

	app.Before = func(c *cli.Context) error {
		kubeconfig, namespace := c.GlobalString("config"), c.GlobalString("namespace")

		clientset, err := kubernetes.CreateKubeClient(kubeconfig)
		if err != nil {
			logrus.Errorf("failed to create kube client: %v", err)
			return err
		}

		kubernetes.KubeClient = clientset

		if _, err = kubernetes.KubeClient.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{}); err != nil {
			newNS := &v1.Namespace{}
			newNS.Name = namespace

			_, err := kubernetes.KubeClient.CoreV1().Namespaces().Create(newNS)
			if err != nil {
				logrus.Errorf("failed to create new namespace in the k8s cluster. reason: %v", err)
				return err
			}
		}
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
