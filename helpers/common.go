package helpers

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/pooled"
)

var pool = pooled.NewMap(45*time.Second, 400*time.Millisecond, getConnection)

func getConnection(key interface{}) (pooled.Conn, error) {
	address, ok := key.(string)
	if !ok {
		return nil, fmt.Errorf("key must be a string")
	}

	conn, err := net.DialTimeout("tcp", address+":3629", 10*time.Second)
	if err != nil {
		return nil, err
	}

	// read the NOKEY line
	pconn := pooled.Wrap(conn)

	//sending "ESC/VP.net" in order to allow other commands
	cmd := []byte{0x45, 0x53, 0x43, 0x2F, 0x56, 0x50, 0x2E, 0x6E, 0x65, 0x74, 0x10, 0x03, 0x00, 0x00, 0x00, 0x00}

	s, err := writeAndRead(pconn, cmd, 5*time.Second, ' ')
	if err != nil {
		log.L.Warnf("there was an error sending the ESC/VP.net command: %v", err)
		return nil, err
	}
	log.L.Infof("connection string: %s\n", s)

	return pconn, nil
}

func writeAndRead(conn pooled.Conn, cmd []byte, timeout time.Duration, delim byte) (string, error) {
	conn.SetWriteDeadline(time.Now().Add(timeout))

	n, err := conn.Write(cmd)
	switch {
	case err != nil:
		return "", err
	case n != len(cmd):
		return "", fmt.Errorf("wrote %v/%v bytes of command 0x%x", n, len(cmd), cmd)
	}

	b, err := conn.ReadUntil(delim, timeout)
	if err != nil {
		return "", err
	}

	conn.Log().Debugf("Response from command: 0x%x", b)
	return strings.TrimSpace(string(b)), nil
}
