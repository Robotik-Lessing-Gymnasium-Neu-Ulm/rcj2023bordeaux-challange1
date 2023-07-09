package teensy

import (
	"fmt"
	"log"
	"unsafe"

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

func SendInt(number int, port serial.Port) {
	_, err := port.Write(IntToByteArray(int64(number)))
	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Fatalf("Error sending data to teensy: %s\n", err)
	}
	fmt.Printf("Send succesfully: %d\n", number)
}

func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[size-1-i] = byt
	}
	return arr
}
