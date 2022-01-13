package device

import "errors"

type Device struct {
	ID                            uint64
	ResourceCapacity              uint
	capacityConfigurations        []CapacityConfiguration
	optimalCapacityConfigurations []CapacityConfiguration
	capacityAlgorithm             IDeviceCapacityAlgorithm
}

func (d *Device) GetOptimalCapacityConfigurations() []CapacityConfiguration {
	return d.optimalCapacityConfigurations
}

func (d *Device) GetCapacityConfigurations() []CapacityConfiguration {
	return d.capacityConfigurations
}

func (d *Device) CalculateCapacityConfigurations(fgTasks []Task, bgTasks []Task) (error, []CapacityConfiguration) {
	if d.capacityAlgorithm == nil {
		return errors.New("No algorithm selected"), nil
	}
	d.capacityConfigurations = d.capacityAlgorithm.CalculateCapacity(d, fgTasks, bgTasks)
	return nil, d.capacityConfigurations
}

func (d *Device) CalculateOptimalConfigurations(fgTasks []Task, bgTasks []Task) (error, []CapacityConfiguration) {
	if d.capacityAlgorithm == nil {
		return errors.New("No algorithm selected"), nil
	}

	d.optimalCapacityConfigurations = d.capacityAlgorithm.CalculateOptimalCapacity(d, fgTasks, bgTasks)
	return nil, d.optimalCapacityConfigurations
}


func CreateNewDevice(id uint64, capacity uint, algorithmType CapacityAlgorithmType) *Device {
	return &Device{
		ID:                            id,
		ResourceCapacity:              capacity,
		capacityConfigurations:        make([]CapacityConfiguration, 0),
		optimalCapacityConfigurations: make([]CapacityConfiguration, 0),
		capacityAlgorithm:             selectCapacityAlgorithm(algorithmType),
	}
}

func selectCapacityAlgorithm(algorithmType CapacityAlgorithmType) *SimpleDeviceCapacityAlgorithm {

	switch algorithmType {
	case SimpleAlgorithm:
		return &SimpleDeviceCapacityAlgorithm{}
	default:
		return &SimpleDeviceCapacityAlgorithm{}
	}

}