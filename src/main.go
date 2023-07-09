package main

import (
	"rcj2023bordeau-challange1/src/teensy"
)

func main() {
	teensy.GetAllPorts()
	PORT := teensy.Communication("COM5")
	teensy.SendInt(69, PORT)

	//gui.InitGui()
	//vision.Init()
}
