package gosudoku

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type TCPConnection struct {
	conn net.Conn
	addr string
	port int
}

var boxManager TCPConnection
var boxConnections map[string]TCPConnection

// Establish initial connection to box manager
func ConnectToManager(maddress *string, mport *int, lport *int) {
	boxManager.addr = *maddress
	boxManager.port = *mport
	boxManager.connect()
	reply := boxManager.sendMessage(myBox.id + "," + getLocalIP().String() + "," + strconv.Itoa(*lport))
	if reply == "OK" {
		log.Println("connected to boxmanager.")
		connectToBoxes()
		sendInitialConfig()
	} else {
		panic(errors.New("connection to boxmanager failed"))
	}
}

// Connect to corresponding Boxes based on myBox
func connectToBoxes() {
	for _, boxID := range boxMap[myBox.id] {
		reply := boxManager.sendMessage(boxID)
		if checkIP(&reply) {
			log.Println("IP address of " + boxID + " is: " + reply)
			addr := strings.Split(reply, ",")
			port, err := strconv.Atoi(addr[1])
			if err != nil {
				panic(err)
			}
			boxConnections[boxID] = connectToBox(addr[0], port)
		} else {
			log.Println("malformed ip address")
		}
	}
}

// Connect to Box
func connectToBox(addr string, port int) TCPConnection {
	var connection TCPConnection
	connection.addr = addr
	connection.port = port
	connection.connect()
	return connection
}

// Send initial config to all connected boxes
func sendInitialConfig() {
	for _, box := range boxConnections {
		for key, val := range myBox.values {
			x, y := getCoordinatesForIndex(key)
			box.sendMessage(myBox.id + "," + strconv.Itoa(x) + "," + strconv.Itoa(y) + ":" + strconv.Itoa(val))
		}
	}
}

// Sends message with value to all neighbors
func sendToNeighbors(x, y, val int) {
	for _, neighbor := range boxConnections {
		neighbor.sendMessage(myBox.id + "," + strconv.Itoa(x) + "," + strconv.Itoa(y) + ":" + strconv.Itoa(val))
	}
}

// Check IP/Port answer from boxManager
func checkIP(ip *string) bool {
	r, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]),[0-9]+$`)
	return r.MatchString(*ip)
}

// Connect to a TCP-Server
func (t *TCPConnection) connect() {
	var err error
	t.conn, err = net.Dial("tcp", t.addr+":"+strconv.Itoa(t.port))
	if err != nil {
		panic(err)
	}
}

// Send message via TCP
func (t *TCPConnection) sendMessage(message string) string {
	_, err := fmt.Fprintf(t.conn, message+"\n")
	checkErr(err)

	reply, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(reply, "\n")
}

// Launching a TCP Server on given port number.
// It handles all incoming request from other boxConnections
func LaunchTCPServer(port *int) {
	go func() {
		log.Println("Launching TCP Server")

		ln, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
		defer ln.Close()

		checkErr(err)

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Error accepting: ", err.Error())
			}
			go handleTCPRequest(conn)
		}
	}()
}

// Handle TCP requests from box manager
func handleTCPRequest(conn net.Conn) {
	// Will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	message = strings.TrimRight(message, "\n")

	if checkMessageFormat(message) {
		r := regexp.MustCompile(`^(BOX_[A,D,G])[1,4,7],([0-2]),([0-2]):([1-9])$`)
		matches := r.FindStringSubmatch(message)
		if strContains(boxMap[myBox.id], matches[1]) {
			val, err := strconv.Atoi(matches[4])
			checkSoftErr(err)
			if matches[1] == myBox.id[:len(myBox.id)-1] {
				y, err := strconv.Atoi(matches[3])
				checkSoftErr(err)
				myBox.setColValue(y, val)
			} else {
				x, err := strconv.Atoi(matches[2])
				checkSoftErr(err)
				myBox.setRowValue(x, val)
			}
		} else {
			log.Println("ALERT: STRANGER DANGER!!!")
		}
	}
	err := conn.Close()
	checkErr(err)
}
