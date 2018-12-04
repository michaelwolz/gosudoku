package gosudoku

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type TCPConnection struct {
	conn net.Conn
	addr string
	port int
}

var BoxManager TCPConnection
var boxes map[string]TCPConnection

// Establish initial connection to box manager
func ConnectToManager(maddress *string, mport *int, lport *int) {
	BoxManager.addr = *maddress
	BoxManager.port = *mport
	BoxManager.connect()
	reply := BoxManager.sendMessage(MyBox.id + "," + getLocalIP().String() + "," + strconv.Itoa(*lport))
	log.Println(reply)

	// Testing only
	//result := BoxManager.sendMessage("GET(BOX_B1)")
	//fmt.Println(result)
}

// Connect to Box
func ConnectToBox(addr *string, port *int) net.Conn {
	return nil
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
	return reply
}

// Launching a TCP Server on given port number.
// It handles all incoming request from other boxes
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
				fmt.Println("Error accepting: ", err.Error())
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
		conn.Write([]byte("getRow\n"))
	case "getCol":
		conn.Write([]byte("getCol\n"))
	default:
		log.Println("Unknown Command received!")
		conn.Write([]byte("Unknown Command\n"))
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
