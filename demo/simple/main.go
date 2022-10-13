package main

import (
	"fmt"
	"log"
	"serial"
	"strings"
	"time"
)

func main() {
	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open("COM1", mode)
	if err != nil {
		log.Fatal(err)
	}

	// Send the string "10,20,30\n\r" to the serial port
	n, err := port.Write([]byte("10,20,30\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	// Read and print the response

	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				n, err := port.Write([]byte("10,20,30\n\r"))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Sent %v bytes\n", n)
				fmt.Println("Tick at", t)
			}
		}
	}()

	buff := make([]byte, 100)
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		fmt.Printf("%s", string(buff[:n]))

		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			break
		}
	}

	for {
		time.Sleep(1024 * time.Second)
	}
}
