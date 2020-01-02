package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	roomNum := os.Getenv("RoomNum")
	machineType := os.Getenv("MachineType")

	requestBody, err := json.Marshal(map[string]interface{}{
		"roomnum":     roomNum,
		"machinetype": machineType,
		"status":      -1,
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
