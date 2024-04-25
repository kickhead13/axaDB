package axa_server

import (
	"fmt"
	"net"
	"strconv"
	"bufio"
	"io"
)

import (
	"axaDB/src/dberrs"
)

func testStart(ip string, port int) {
	ln, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("(axa server): listening on " + ip + ":" + strconv.Itoa(port) + "...")
	
	for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
		fmt.Printf("(axa server): new connection: (%s)\n", conn.RemoteAddr())
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {

    defer conn.Close()
	wrapped_conn := bufio.NewReader(conn)
	buff := make([]byte, 256)
	remoteAddr := fmt.Sprintf("%s", conn.RemoteAddr())
	for {
    	m_len, err := wrapped_conn.ReadByte()
    	if err != nil {
			conn.Close()
			break
    	}

		_, err = io.ReadFull(wrapped_conn, buff[:int(m_len)])
		if err != nil {
			conn.Close()
			return
		}

    	fmt.Printf("(axa server) received [ %s ] from %s \n", buff[:int(m_len)], remoteAddr)
	}
	fmt.Printf("(axa server): connection to %s closed...\n", remoteAddr )
}

func Start(_args []string) dberrs.AxaErr{
	port := 13131
	ip := "127.0.0.1"
	testStart(ip, port)
	return dberrs.DB_NORM()
}

func Halt() dberrs.AxaErr{
	return dberrs.DB_NORM()
} 

