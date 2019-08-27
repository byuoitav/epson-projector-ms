package helpers

import (
	"net"
	"time"

	"github.com/byuoitav/common/pooled"

	"github.com/byuoitav/common/log"
)

func getConnection(address string) pooled.Conn {
	address += ":3629"

	netConn, err := net.Dial("tcp", address)
	if err != nil {
		log.L.Debugf("There was a problem opening the connection: %v", err)
		return nil
	}

	conn := pooled.Wrap(netConn)

	//sending "ESC/VP.net" in order to allow other commands
	cmd := []byte{0x45, 0x53, 0x43, 0x2F, 0x56, 0x50, 0x2E, 0x6E, 0x65, 0x74, 0x10, 0x03, 0x00, 0x00, 0x00, 0x00}

	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		log.L.Warnf("There was an error sending the ESC/VP.net command: %v", err)
		return nil
	case n != len(cmd):
		log.L.Warnf("only sent %v/%v bytes\n", n, len(cmd))
		return nil
	}

	bytes, err := conn.ReadUntil(' ', 5*time.Second)
	if err != nil {
		log.L.Warnf("There was an error after sending the ESC/VP.net command: %v", err)
		return nil
	}
	log.L.Infof("Bytes returned: %s", bytes)
	return conn
}
