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
	fmt.Fprintf(t.conn, message+"\n")

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

		if err != nil {
			panic(err)
		}

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

	// Handle command
	switch string(message) {
	case "getRow":
		conn.Write([]byte("sendVal\n"))
	default:
		log.Println("unknown command received!")
		conn.Write([]byte("unknown command\n"))
	}

	conn.Close()
}

// Get Local IP Address (https://gist.github.com/jniltinho/9787946)
func getLocalIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}
