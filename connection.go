package gosudoku

import (
	"bufio"
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TCPConnection struct {
	conn net.Conn
	addr string
	port int
	name string
}

var boxManager TCPConnection
var boxConnections = make(map[string]TCPConnection)
var sendMutex = &sync.Mutex{}
var Done = make(chan int)

// Establish initial connection to box manager
func ConnectToManager(maddress *string, mport *int, lport *int) {
	boxManager.addr = *maddress
	boxManager.port = *mport
	boxManager.connect()
	reply := boxManager.sendMessage(myBox.id+","+getLocalIP().String()+","+strconv.Itoa(*lport), true)
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
	for _, boxID := range neighbors[myBox.id] {
		reply := boxManager.sendMessage(boxID, true)
		if checkIP(&reply) {
			addr := strings.Split(reply, ",")
			port, err := strconv.Atoi(addr[1])
			if err != nil {
				panic(err)
			}
			boxConnections[boxID] = connectToBox(addr[0], port, boxID)
		} else {
			log.Println("malformed ip address")
		}
	}
}

// Connect to Box
func connectToBox(addr string, port int, name string) TCPConnection {
	var connection TCPConnection
	connection.addr = addr
	connection.port = port
	connection.name = name
	connection.connect()
	return connection
}

// Send initial config to all connected boxes
func sendInitialConfig() {
	for _, box := range boxConnections {
		for key, val := range myBox.values {
			if val != 0 {
				x, y := getCoordinatesForIndex(key)
				box.sendMessage(myBox.id+","+strconv.Itoa(x)+","+strconv.Itoa(y)+":"+strconv.Itoa(val), false)
			}
		}
	}
}

// Sends message with value to all neighbors
func sendToNeighbors(x, y, val int) {
	for _, neighbor := range boxConnections {
		neighbor.sendMessage(myBox.id+","+strconv.Itoa(x)+","+strconv.Itoa(y)+":"+strconv.Itoa(val), false)
	}
}

// Redirects an incoming message to other boxes which are influenced by the sender e.g. A1 -> G7
func redirectToNeighbor(message string, sender string) {
	if _, ok := redirectNeighbors[myBox.id][sender]; ok {
		for _, conn := range boxConnections {
			if conn.name == redirectNeighbors[myBox.id][sender] {
				conn.sendMessage(message, false)
			}
		}
	}
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
func (t *TCPConnection) sendMessage(message string, expectReply bool) string {
	sendMutex.Lock() // Locking send message, so that two messages aren't send at the same time
	_, err := t.conn.Write([]byte(message + "\n"))
	checkErr(err)

	var reply string
	if expectReply {
		// TODO: Add timeout!
		reply, err = bufio.NewReader(t.conn).ReadString('\n')
		reply = strings.TrimSuffix(reply, "\n")
		checkErr(err)
	}

	time.Sleep(10 * time.Millisecond)
	sendMutex.Unlock()
	return reply
}

// Launching a TCP Server on given port number.
// It handles all incoming request from other boxConnections
func LaunchTCPServer(port *int) {
	go processMessages() // Starts worker for processing incoming messages
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

// Handle TCP requests from box manager and put incoming messages in the message queue
func handleTCPRequest(conn net.Conn) {
	var message string

	for {
		message, _ = bufio.NewReader(conn).ReadString('\n')
		msgChan <- message
		conn.Write([]byte("OK\n"))
	}
}
