package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

type Machine struct {
	RoomNum     string
	MachineType string
	Status      int
}

func intializePins() rpio.Pin {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	pin := rpio.Pin(19)
	pin.Mode(rpio.Pwm)
	pin.Freq(9000)
	pin.Input()
	return pin
	//Need to verify that this works with sensors that I have
	//Need to verify the freq rate is also correct

}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func readSensorData(pin rpio.Pin) int {
	res := pin.Read()
	//Add logic in once I see what the output usually is.
	//Update status in Struct
}

func sendRequest(m Machine) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"roomnum":     m.RoomNum,
		"machinetype": m.MachineType,
		"status":      m.Status,
	})
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post("http://localhost:8081/statusChange", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	defer resp.Body.Close()

}
func main() {

	roomNum := os.Getenv("RoomNum")
	machineType := os.Getenv("MachineType")
	var m Machine
	m.RoomNum = roomNum
	m.MachineType = machineType
	pin := intializePins()
	for {
		time.Sleep(5 * time.Second)
		x := readSensorData(pin)
		sendRequest(m)
	}
}
