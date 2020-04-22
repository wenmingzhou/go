package utils

import (
	"log"
	"net"
	"strings"
)

func GetOutboundIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
