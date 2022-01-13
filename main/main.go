package main

import (
	"fmt"
	"ikonchallenge/service"
)

var deviceService service.IDeviceService;

func main() {

	deviceService = &service.DeviceService{}

	devices := deviceService.CalculateOptimalCapacityFromFile(".","challenge.in")
	deviceService.WriteOptimalCapacityToFile(".", "challenge.out", devices)

	for _, d:= range devices {
		fmt.Println("DeviceID: ", d.ID)
		for _, cap:= range d.GetOptimalCapacityConfigurations(){
			fmt.Printf("Foreground ID:%d BackgroundID:%d ConfigConsumption:%d\n", cap.ForegroundTask.ID, cap.BackgroundTask.ID, cap.Consumption)
		}
		fmt.Println("")
		
	}
}