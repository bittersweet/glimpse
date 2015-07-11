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

// Channel that is used to send messages to the SSE broker
var notifier chan []byte

const maxPacketSize = 1024 * 1024

// formatMessage strips ANSI codes e.g. `[30m`
func formatMessage(msg []byte) string {
	re := regexp.MustCompile(`\[\d+m`)
	msg = re.ReplaceAllLiteral(msg, []byte(""))

	return string(msg)
}

// Global identifier to generate IDs for sequential messages
// Currently in use to know we have unique objects on the frontend
var messageID = 0

// Process incoming udp payloads and send them as JSON to the frontend via
// Server Sent Events
// We add the current timestamp and a sequential messageID
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

	// Send JSON payload to the SSE broker
	notifier <- output
}

// listenToUDP will receive UDP payloads and process them
func listenToUDP() {
	addr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		log.Fatal("ResolveUDPAddr", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("ListenUDP", err)
	}

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

// Serve 3 static files currently in use, this is all we need right now so
// lets not make it too complicated
func handleIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/":
		http.ServeFile(w, r, "frontend/dist/index.html")
	case "/styles.css":
		http.ServeFile(w, r, "frontend/styles.css")
	case "/app.js":
		http.ServeFile(w, r, "frontend/dist/app.js")
	}
}

// Setup everything
// - Simple HTTP endpoint on port 8001 that responds to SSE connection requests
// - UDP listener that will receive messages
// - HTTP server that serves static files in use for the frontend
func main() {
	broker := NewServer()
	notifier = broker.Notifier

	// Start SSE endpoint that will register connections
	go http.ListenAndServe(":8001", broker)

	// Start listening to UDP payloads
	go listenToUDP()

	http.HandleFunc("/", handleIndex)

	fmt.Println("Will start listening on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Could not start listening on port 8000")
	}
}
