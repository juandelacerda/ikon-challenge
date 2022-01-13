package device

type SimpleDeviceCapacityAlgorithm struct{}

func (a *SimpleDeviceCapacityAlgorithm) CalculateCapacity(d *Device, fgTasks []Task, bgTasks []Task) []CapacityConfiguration {
	capacityConfiguration := make([]CapacityConfiguration, 0)

	for _, fgTask := range fgTasks {
		for _, bgTask := range bgTasks {
			if fgTask.Consumption+bgTask.Consumption <= d.ResourceCapacity {
				capacityConfiguration = append(capacityConfiguration, CapacityConfiguration{
					ForegroundTask: fgTask,
					BackgroundTask: bgTask,
					Consumption:    fgTask.Consumption + bgTask.Consumption,
				})
			}
		}
	}

	return capacityConfiguration
}

func (a *SimpleDeviceCapacityAlgorithm) CalculateOptimalCapacity(d *Device, fgTasks []Task, bgTasks []Task) []CapacityConfiguration {
	capacityConfiguration := make([]CapacityConfiguration, 0)
	optimalConfiguration := make([]CapacityConfiguration, 0)

	capacityConfiguration = a.CalculateCapacity(d, fgTasks, bgTasks)

	for _, config := range capacityConfiguration {
		if config.ForegroundTask.Consumption+config.BackgroundTask.Consumption == d.ResourceCapacity {
			optimalConfiguration = append(optimalConfiguration, config)
		}
	}

	if len(optimalConfiguration) == 0 {
		var nextOptimalConfiguration CapacityConfiguration
		for i, config := range capacityConfiguration {
			if i == 0 {
				nextOptimalConfiguration = config
				continue
			}

			if config.ForegroundTask.Consumption+config.BackgroundTask.Consumption >
				nextOptimalConfiguration.ForegroundTask.Consumption+nextOptimalConfiguration.BackgroundTask.Consumption {
				nextOptimalConfiguration = config
			}
		}

		optimalConfiguration = append(optimalConfiguration, nextOptimalConfiguration)

	}

	return optimalConfiguration
}