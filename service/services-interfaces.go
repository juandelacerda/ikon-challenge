package service

import "ikonchallenge/device"

type IDeviceService interface {
	CalculateOptimalCapacityFromFile(path string, name string)[]*device.Device
	WriteOptimalCapacityToFile(path string, name string, optimalConfiguration []*device.Device)
}
