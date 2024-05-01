package axa_server

import (
	"fmt"
	"net"
	"strconv"
	"bufio"
	"io"
	"sync"
)

import (
	"axaDB/src/dberrs"
)

func testStart(ip string, port int) {
	var execBuffer CritBuffer = InitCritBuffer()
	var responseBuffer CritBuffer = InitCritBuffer()
	var wgroup sync.WaitGroup

	ln, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("(axa server): listening on " + ip + ":" + strconv.Itoa(port) + "...")
	
	wgroup.Add(1)
	go handleExecutioner(&execBuffer, &responseBuffer, &wgroup)

	for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
		fmt.Printf("(axa server): new connection: (%s)\n", conn.RemoteAddr())
		wgroup.Add(1)
        go handleConnection(conn, &execBuffer, &responseBuffer, &wgroup)
    }
	wgroup.Wait()
}

func handleConnection(conn net.Conn, execBuffer *CritBuffer, responseBuffer *CritBuffer, wgroup *sync.WaitGroup) {

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

		execBuffer.push(fmt.Sprintf("%s", buff[:int(m_len)]), remoteAddr)
    	fmt.Printf("(axa server) received [ %s ] from %s \n", buff[:int(m_len)], remoteAddr)

		received := false
		var val string
		for ! received {
			val = responseBuffer.readValueOfKey(remoteAddr)
			if val != "" {

				fmt.Printf("(axa server) received client %s's response:\n * %s\n", remoteAddr, val)

				received = true
				m_len := byte(len(val))
				m_buff := make([]byte, 1)
				m_buff[0] = m_len

				_, err := conn.Write(append(m_buff[:], []byte(val)...))
				if err != nil {
					fmt.Println(err)
					return 
				}	
			}
		}
		
	}
	fmt.Printf("(axa server): connection to %s closed...\n", remoteAddr )
	wgroup.Done()
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

