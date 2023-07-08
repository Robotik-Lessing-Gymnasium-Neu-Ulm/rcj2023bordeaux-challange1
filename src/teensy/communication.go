package teensy

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func GetAllPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error searching for ports: %s\n", err)
	}

	if len(ports) == 0 {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("No ports avaiable: %s\n", err)
	}

	for _, port := range ports {
		fmt.Printf("Following port found: %v\n", port)
	}
}

func Communication(port string) serial.Port {
	mode := &serial.Mode{
		BaudRate: 115200,
	}

	PORT, err := serial.Open(port, mode)
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error opening USB connection: %s\n", err)
	}

	fmt.Println("Connection successfull")

	return PORT
}
