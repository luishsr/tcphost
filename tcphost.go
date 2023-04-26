package tcphost

import (
	"fmt"
	"log"
	"net"
)

func RunListener(address string, port string, received chan<- string) {

	fmt.Println("New Listener launched at " + address + ":" + port)

	// Create a listener
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Listener returned: %s", err)
	}
	defer l.Close()

	for {
		// Accept new connections
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Unable to accept new connections: %s", err)
		}

		// Create a goroutine that reads and writes-back data
		go func() {
			log.Printf("TCP Session Open")
			// Clean up session when goroutine completes, it's ok to
			// call Close more than once.
			defer c.Close()

			for {
				b := make([]byte, 1024)

				// Read from TCP Buffer
				_, err := c.Read(b)
				if err != nil {
					log.Printf("Error reading TCP Session: %s", err)
					break
				}

				//Add Read data to the channel
				//TO-DO: continue implementation
				received <- string(b)

				// Write-back data to TCP Client
				_, err = c.Write(b)
				if err != nil {
					log.Printf("Error writing TCP Session: %s", err)
					break
				}
			}
		}()

	}

}

func RunSender(address string, port string) net.Conn {

	fmt.Println("New Sender launched at " + address + ":" + port)
	return
}
