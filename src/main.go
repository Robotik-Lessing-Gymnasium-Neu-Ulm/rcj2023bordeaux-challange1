package main

import (
	"rcj2023bordeau-challange1/src/gui"
	"rcj2023bordeau-challange1/src/teensy"
)

func main() {
	teensy.GetAllPorts()
	_ = teensy.Communication("COM5")

	gui.InitGui()
}
