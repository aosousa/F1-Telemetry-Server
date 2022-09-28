package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"

	h "github.com/aosousa/F1-Telemetry-Server/handlers"
	"github.com/aosousa/F1-Telemetry-Server/logger"
	"github.com/aosousa/F1-Telemetry-Server/models"
)

const version = "0.0.1"

func init() {
	// set up Config struct before setting up the server
	fmt.Println("Configuration file: Loading")
	h.InitConfig()
	fmt.Printf("Configuration file: OK\n\n")

	// initialize database through Config struct
	// fmt.Println("System database: Checking")
	// h.InitDatabase()
	// fmt.Printf("System database: OK\n\n")
}

func main() {
	port := strconv.Itoa(int(h.Config.UDPPort))
	addr := net.UDPAddr{
		Port: int(h.Config.UDPPort),
		IP:   net.ParseIP("127.0.0.1"),
	}

	connection, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer connection.Close()

	fmt.Printf("F1 Telemetry Server V%s running on port %s\n", version, port)

	buffer := make([]byte, 4096)
	bufferCopy := make([]byte, 4096)
	for {
		_, _, err := connection.ReadFrom(buffer)
		if err != nil {
			logger.Error(err.Error())
		}

		copy(bufferCopy, buffer)

		r := bytes.NewReader(buffer)
		rCopy := bytes.NewReader(bufferCopy)

		var test models.Header
		if err := binary.Read(r, binary.LittleEndian, &test); err != nil {
			fmt.Println("binary.Read failed:", err)
		}

		// fmt.Printf("%+v\n", test)

		if test.PacketID == 1 {
			var sessionData models.PacketSessionData

			if err := binary.Read(rCopy, binary.LittleEndian, &sessionData); err != nil {
				fmt.Println("binary.Read failed:", err)
			}

			fmt.Print("session data")
			fmt.Printf("%+v\n", sessionData)
		}
	}
}
