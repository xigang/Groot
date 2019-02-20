package cmd

import "github.com/urfave/cli"

type submitTFJob struct {
	//ps
	PSReplicas int
	PSPort     int
	PSImage    string
	PSCpu      string
	PSMemory   string

	//worker
	WorkerReplicas int
	WorkePort      int
	WorkerImage    string
	WorkerCpu      string
	WorkerMemory   string
	WorkerGpus     int
	WorkerDir      string

	CleanPodPolicy string //eg: Running, All, None

	//Chief
	UseChief      bool
	ChiefReplicas int
	ChiefPort     int
	ChiefImage    string
	ChiefCpu      string
	ChiefMemory   string

	//Evaluator
	UseEvaluator      bool
	EvaluatorReplicas int
	EvaluatorPort     int
	EvaluatorImage    string
	EvaluatorCpu      string
	EvaluatorMemory   string
}

var TFJobSubCommand = cli.Command{
	Name:    "tfjob",
	Aliases: []string{"tf"},
	Usage:   "submit a TFjob as training job.",
	Action: func(c *cli.Context) error {
		return nil
	},
}
