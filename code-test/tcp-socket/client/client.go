package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8081")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	for {
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\n")
		if str == "quit" {
			break
		}
		n, _ := conn.Write([]byte(str))
		fmt.Printf("wrote %d bytes to %s\n", n, tcpAddr.String())
	}

}
