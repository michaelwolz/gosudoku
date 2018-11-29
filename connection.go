package gosudoku

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Connection interface {
	getRow() []int
	getCol() []int
}

type TCPConnection struct {
	conn net.Conn
	addr string
	port int
}

// List of addresses to corresponding box numbers
var boxList = make([]string, 9)

// Use mDNS to find other boxes
func FindBoxes(boxID *int, port *int) {
	go registerService(boxID, port)
	searchForClients()
}

// Launching a TCP Server on given port number.
// It handles all incoming request from other boxes
func launchTCPServer(port int) {
	log.Println("Launching TCP Server")

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer ln.Close()

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleTCPRequest(conn)
	}
}

// Handle TCP requests from other boxes
func handleTCPRequest(conn net.Conn) {
	// Will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')

	// Handle command
	switch string(message) {
	case "getRow":
		conn.Write([]byte("getRow\n"))
	case "getCol":
		conn.Write([]byte("getCol\n"))
	default:
		log.Println("Unknown Command received!")
		conn.Write([]byte("Unknown Command\n"))
	}

	conn.Close()
}

// Connect to a TCP-Server
func (t *TCPConnection) tcpconnect() net.Conn {
	conn, err := net.Dial("tcp", t.addr+":"+strconv.Itoa(t.port))
	if err != nil {
		log.Println(err)
	}
	return conn
}

// Send message via TCP
func (t *TCPConnection) sendTCPMessage(message string) {
	fmt.Fprintf(t.conn, message+"\n")

	reply, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	fmt.Print("Reply from server: " + reply)
}
