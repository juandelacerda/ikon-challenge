package service

import (
	"errors"
	"fmt"
	"ikonchallenge/device"
	"strconv"
	"strings"
)

const LinesPerFormatDevice int = 3


func parseTupleToTaskFormat(lines []string) (error, uint, []device.Task, []device.Task) {
	if !isValidTupleFormat(lines){
		return errors.New("Tuple format is invalid"), 0, nil, nil
	}
	capacity, err := strconv.Atoi(strings.Trim(lines[0], "\r"))

	if(err !=nil){
		fmt.Println(err)
		return errors.New("Unable parse device capacity"), 0, nil, nil
	}

	fgTaskTuples := strings.Split(lines[1], "), ") 
	bgTaskTuples := strings.Split(lines[2], "), ")

	fgTasks :=  tuplesToTask(fgTaskTuples)
	bgTasks :=  tuplesToTask(bgTaskTuples)

	return nil, uint(capacity), fgTasks, bgTasks
}

func tuplesToTask(tuplesTasks []string) []device.Task{
	tasks:= make([]device.Task, 0)
	for _,tuple := range tuplesTasks {
		tuple := strings.ReplaceAll(tuple, "(", "")
		tuple = strings.ReplaceAll(tuple, ")", "")
		tuple = strings.Trim(tuple, "\r")
		tupleData := strings.Split(tuple, ",")
		id,_:= strconv.Atoi(tupleData[0])		

		consumption, err := strconv.Atoi(strings.Trim(tupleData[1], " "))
		if err!=nil {
			fmt.Println(err)
		}
		tasks = append(tasks, device.Task{
			ID: uint64(id),
			Consumption: uint(consumption),
		})
	}

	return tasks
}

func isValidTupleFormat(lines []string) bool {
	numberOfLines := len(lines)
	if numberOfLines%3!=0 {
		return false
	}
	return true
}

func generateRandomDeviceID() uint64{
	return 123
}

func parseDeviceCapacityConfigurationToTuples(d *device.Device) string{
	tuples:= make([]string, 0)

	for _, config:= range d.GetOptimalCapacityConfigurations() {
		tuples = append(tuples, capacityConfigurationToTuple(config))
	} 

	return  strings.Join(tuples, ", ") ;
}

func capacityConfigurationToTuple(c device.CapacityConfiguration) string {
	return "(" + strconv.FormatUint(c.ForegroundTask.ID, 10) + ", "+ strconv.FormatUint(c.BackgroundTask.ID, 10) +")"
}