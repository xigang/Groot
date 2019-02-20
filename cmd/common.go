package cmd

type commonArgs struct {
	TrainName      string //train name
	TrainMode      string //single or dist
	TrainNamespace string //the namespace of the train job
}
