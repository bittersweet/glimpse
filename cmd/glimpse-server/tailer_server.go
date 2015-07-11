package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var notifier chan []byte

const maxPacketSize = 1024 * 1024

func formatMessage(msg []byte) string {
	re := regexp.MustCompile(`\[\d+m`)
	msg = re.ReplaceAllLiteral(msg, []byte(""))

	return string(msg)
}

var messageID = 0

func processMessage(msg []byte) {
	messageID++
	fmt.Printf("%s", msg)
	formattedMsg := formatMessage([]byte(msg))

	t := time.Now()
	// Format based on default reference time "Mon Jan 2 15:04:05 MST 2006"
	timestamp := t.Format("02/01/06 15:04:05")

	event := map[string]string{
		"id":        strconv.Itoa(messageID),
		"timestamp": timestamp,
		"line":      formattedMsg,
	}

	output, err := json.Marshal(event)
	if err != nil {
		log.Fatal("MarshalIndent", err)
	}

	fmt.Println("output:", string(output))

	notifier <- output
}

func listenToUDP(conn *net.UDPConn) {
	buffer := make([]byte, maxPacketSize)
	for {
		bytes, err := conn.Read(buffer)
		if err != nil {
			log.Println("UDP read error: ", err.Error())
			continue
		}

		msg := make([]byte, bytes)
		copy(msg, buffer)
		processMessage(msg)
	}
}

func main() {
	broker := NewServer()
	notifier = broker.Notifier
	go http.ListenAndServe(":8001", broker)

	addr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		log.Fatal("ResolveUDPAddr", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("ListenUDP", err)
	}
	listenToUDP(conn)
}
