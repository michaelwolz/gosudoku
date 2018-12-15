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
var boxConnections = make(map[string]TCPConnection)

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
	// TODO: Close connection to box manager
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
			if val != 0 {
				x, y := getCoordinatesForIndex(key)
				reply := box.sendMessage(myBox.id + "," + strconv.Itoa(x) + "," + strconv.Itoa(y) + ":" + strconv.Itoa(val))
				fmt.Println(reply)
			}
		}
	}
}

// Sends message with value to all neighbors
func sendToNeighbors(x, y, val int) {
	for _, neighbor := range boxConnections {
		reply := neighbor.sendMessage(myBox.id + "," + strconv.Itoa(x) + "," + strconv.Itoa(y) + ":" + strconv.Itoa(val))
		fmt.Println(reply)
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
	_, err := t.conn.Write([]byte(message + "\n"))
	checkErr(err)

	reply, err := bufio.NewReader(t.conn).ReadString('\n')
	checkErr(err)

	return strings.TrimSuffix(reply, "\n")
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

// TODO: unblock connection
// Handle TCP requests from box manager
func handleTCPRequest(conn net.Conn) {
	var message string
	var err error

	// Will listen for message to process ending in newline (\n)
	message, err = bufio.NewReader(conn).ReadString('\n')
	message = strings.TrimSuffix(message, "\n")
	_, err = conn.Write([]byte("Message received."))
	checkErr(err)

	if checkMessageFormat(message) {
		r := regexp.MustCompile(`^(BOX_[A,D,G][1,4,7]),([0-2]),([0-2]):([1-9])$`)
		matches := r.FindStringSubmatch(message)
		if strContains(boxMap[myBox.id], matches[1]) {
			fmt.Println("Value of message: " + message + " is: " + matches[4])
			val, err := strconv.Atoi(matches[4])
			fmt.Println(val)
			checkSoftErr(err)
			if matches[1][:len(matches[1])-1] == myBox.id[:len(myBox.id)-1] {
				x, err := strconv.Atoi(matches[2])
				checkSoftErr(err)
				myBox.setColValue(x, val)
			} else {
				y, err := strconv.Atoi(matches[3])
				checkSoftErr(err)
				myBox.setRowValue(y, val)
			}
		} else {
			log.Println("ALERT: STRANGER DANGER!!!: " + myBox.id + " -> " + matches[1])
		}
	}
}
