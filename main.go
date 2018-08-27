package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var highTemp = 60000
var lowTemp = 40000

func main() {
	tempFile, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	tempString := make([]byte, 5)
	tempString, err = ioutil.ReadAll(tempFile)
	if err != nil {
		panic(err)
	}

	temp, err := strconv.Atoi(strings.Trim(string(tempString), "\n"))
	if err != nil {
		panic(err)
	}

	raspiAdaptor := raspi.NewAdaptor()
	fun := gpio.NewFunDriver(raspiAdaptor, "18")

	work := func() {
		gobot.Every(1*time.Second, func() {
			fun.Toggle()
		})
	}

	robot := gobot.NewRobot("bot", []gobot.Connection{raspiAdaptor}, []gobot.Device{fun}, work)

	robot.Start()

	if temp >= highTemp {
		fmt.Print("high")
	} else if temp >= lowTemp {
		fmt.Print("middle")
	} else {
		fmt.Print("log")
	}
}
