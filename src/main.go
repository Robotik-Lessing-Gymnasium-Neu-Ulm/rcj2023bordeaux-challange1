package main

import (
	"rcj2023bordeau-challange1/src/teensy"
	"rcj2023bordeau-challange1/src/vision"
)

func main() {
	teensy.GetAllPorts()
	PORT := teensy.Communication("COM5")
	teensy.SendInt(69, PORT)

	//gui.InitGui()
	vision.Init()
}
