package device

type Task struct {
	ID          uint64
	Consumption uint
}

type CapacityConfiguration struct {
	ForegroundTask Task
	BackgroundTask Task
	Consumption    uint
}

type CapacityAlgorithmType int

const (
	SimpleAlgorithm CapacityAlgorithmType = iota
	OtherAlgorithm
)
