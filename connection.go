package gosudoku

import (
	"net"
)

type TCPConnection struct {
	conn net.Conn
	addr string
	port int
	name string
}

var boxManager TCPConnection
var Done = make(chan int)

/*// Establish initial connection to box manager
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
*/
