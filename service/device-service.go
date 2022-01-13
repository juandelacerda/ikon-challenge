package service

import (
	"fmt"
	"ikonchallenge/device"
	"ikonchallenge/file"
)

type DeviceService struct{}


func (s *DeviceService) CalculateOptimalCapacityFromFile(path string, filename string) []*device.Device {
	
	devices:= make([]*device.Device, 0)
	
	_,lines := file.ReadLines(path, filename)

	numberOfDevices := int (len(lines)/LinesPerFormatDevice)
	for i:= 0; i<numberOfDevices; i++ {
		err, deviceCapacity, fgTasks, bgTasks := parseTupleToTaskFormat(lines[i*3: i*3 + LinesPerFormatDevice])
		if(err != nil){
			fmt.Println(err)
		}
		d:= device.CreateNewDevice(generateRandomDeviceID(), deviceCapacity, device.SimpleAlgorithm)
		d.CalculateOptimalConfigurations(fgTasks, bgTasks)
		devices = append(devices, d)
	}

	return devices
}

func (s *DeviceService) WriteOptimalCapacityToFile(path string, filename string, devices []*device.Device){
	lines:= make([]string, 0)

	for _, d := range devices {
		lines = append(lines, parseDeviceCapacityConfigurationToTuples(d))
	}

	file.WriteLines(path, filename, lines)
}

func CreateNewDeviceService() *DeviceService{
	return &DeviceService{}
}

