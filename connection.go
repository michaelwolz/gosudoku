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

func ConnectToManager(maddress *string, mport *int, lport *int) {
	var connection TCPConnection
	connection.addr = *maddress
	connection.port = *mport
	connection.connect()
	connection.sendMessage(MyBox.id + "," + getLocalIP().String() + "," + strconv.Itoa(*lport))
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
func (t *TCPConnection) sendMessage(message string) {
	fmt.Fprintf(t.conn, message+"\n")
	reply, err := bufio.NewReader(t.conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print("Reply from server: " + reply)
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
