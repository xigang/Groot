package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type TFJob struct {
	//ps
	PSReplicas int    `json:"ps"`
	PSPort     int    `json:"ps-port"`
	PSImage    string `json:"ps-image"`
	PSCpu      string `json:"ps-cpu"`
	PSMemory   string `json:"ps-memory"`

	//worker
	WorkerReplicas int    `json:"workers"`
	WorkePort      int    `json:"worker-port"`
	WorkerImage    string `json:"worker-image"`
	WorkerCpu      string `json:"worker-cpu"`
	WorkerMemory   string `json:"worker-memory"`
	WorkerGpus     int    `json:"num-gpus"`
	WorkerDir      string `json:"worker-dir"`

	CleanPodPolicy string `json:"clean-pod-policy"` //eg: Running, All, None

	//Chief
	UseChief      bool   `json:"use-chief"`
	ChiefReplicas int    `json:"chiefs"`
	ChiefPort     int    `json:"chief-port"`
	ChiefImage    string `json:"chief-image"`
	ChiefCpu      string `json:"chief-cpu"`
	ChiefMemory   string `json:"chief-memory"`

	//Evaluator
	UseEvaluator      bool   `json:"use-evaluator"`
	EvaluatorReplicas int    `json:"evaluators"`
	EvaluatorPort     int    `json:"evaluator-port"`
	EvaluatorImage    string `json:"evaluator-image"`
	EvaluatorCpu      string `json:"evaluator-cpu"`
	EvaluatorMemory   string `json:"evaluator-memory"`

	//common args
	commonArgs `json:"common-args"`
}

var TFJobSubCommand = cli.Command{
	Name:    "tfjob",
	Aliases: []string{"tf"},
	Usage:   "submit a TFjob as training job.",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "",
			Usage: "training job name.",
		},
		cli.StringFlag{
			Name:  "namespace",
			Value: "default",
			Usage: "the namespace of the job.",
		},
		cli.IntFlag{
			Name:  "ps",
			Value: 1,
			Usage: "the number of the parameter server.",
		},
		cli.IntFlag{
			Name:  "ps-port",
			Value: 2222,
			Usage: "the port of the paramter server.",
		},
		cli.StringFlag{
			Name:  "ps-image",
			Value: "",
			Usage: "the docker image for tensorflow workers.",
		},
		cli.IntFlag{
			Name:  "ps-cpu",
			Usage: "the cpu resource to use for the parameter servers, eg: 1 for 1 core."},
		cli.StringFlag{
			Name:  "ps-memory",
			Value: "",
			Usage: "the memory resource to use for the parameter servers, eg: 1Gi.",
		},
		cli.IntFlag{
			Name:  "workers",
			Value: 1,
			Usage: "the worker number to run the distributed training.",
		},
		cli.IntFlag{
			Name:  "worker-port",
			Value: 2222,
			Usage: "the port of the worker.",
		},
		cli.StringFlag{
			Name:  "worker-image",
			Value: "",
			Usage: "the dokcer image for tensorflow worker.",
		},
		cli.IntFlag{
			Name:  "worker-cpu",
			Usage: "the cpu resource to use for worker, eg 1 for 1 core.",
		},
		cli.StringFlag{
			Name:  "worker-memory",
			Value: "",
			Usage: "the memory resource to use for the worker. eg: 1Gi.",
		},
		cli.IntFlag{
			Name:  "num-gpus",
			Usage: "the nvidia gpu resource to use for the worker. eg: 1 for 1 gpu device.",
		},
		cli.StringFlag{
			Name:  "worker-dir",
			Value: "/root",
			Usage: "working directory to extract the code. If using syncMode, the $workingDir/code contains the code.",
		},
		cli.StringFlag{
			Name:  "device",
			Value: "cpu",
			Usage: "the device to use for training job.",
		},
	},
	Action: submitTFJob,
}

var submitTFJob = func(c *cli.Context) error {
	tfjob := NewTFJob(c)

	err := tfjob.CreateTFJob()
	if err != nil {
		return err
	}

	return nil
}

func (tf *TFJob) CreateTFJob() error {
	return nil
}

func NewTFJob(c *cli.Context) TFJob {
	checkTFJobFlag(c)

	return TFJob{
		PSReplicas:     c.Int("ps"),
		PSPort:         c.Int("ps-port"),
		PSImage:        c.String("ps-image"),
		PSCpu:          fmt.Sprintf("%d%s", c.Int("ps-cpu")*1000, "m"),
		PSMemory:       c.String("ps-memory"),
		WorkerReplicas: c.Int("workers"),
		WorkePort:      c.Int("worker-port"),
		WorkerImage:    c.String("worker-image"),
		WorkerCpu:      fmt.Sprintf("%d%s", c.Int("worker-cpu")*1000, "m"),
		WorkerMemory:   c.String("worker-memory"),
		WorkerGpus:     c.Int("num-gpus"),
		WorkerDir:      c.String("worker-dir"),
		commonArgs: commonArgs{
			TrainName:      c.String("name"),
			TrainNamespace: c.String("namespace"),
			TrainDevice:    c.String("device"),
		},
	}
}

func checkTFJobFlag(c *cli.Context) {
	trainName := c.String("name")
	if len(trainName) == 0 {
		fmt.Printf("please enter train name.")
		os.Exit(1)
	}

	trainDevice, numGPUs := c.String("device"), c.Int("num-gpus")
	if trainDevice == TRAIN_DEVICE_GPU && numGPUs == 0 {
		fmt.Printf("please enter the number of GPU devices.")
		os.Exit(1)
	}

	psMem := c.String("ps-memory")
	if !strings.Contains(psMem, "Gi") {
		fmt.Printf("please enter the correct ps-memory parameter. eg: 1Gi")
		os.Exit(1)
	}

	workerMem := c.String("worker-memory")
	if !strings.Contains(workerMem, "Gi") {
		fmt.Printf("please enter the correct worker-memory parameter. eg: 1Gi")
		os.Exit(1)
	}
}
