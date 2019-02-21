package cmd

const (
	TRAIN_DEVICE_CPU = "cpu"
	TRAIN_DEVICE_GPU = "gpu"
)

type commonArgs struct {
	TrainName      string `json:"name"`       //train name
	TrainMode      string `json:"train-mode"` //single or dist
	TrainNamespace string `json:"namespace"`  //the namespace of the train job
	TrainDevice    string `json:"device"`
}
