package device

type IDeviceCapacityAlgorithm interface {
	CalculateOptimalCapacity(d *Device, fgTasks []Task, bgTasks []Task) []CapacityConfiguration
	CalculateCapacity(d *Device, fgTasks []Task, bgTasks []Task) []CapacityConfiguration
}
