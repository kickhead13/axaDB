package axa_connect

import (
	"net"
	"fmt"
	"bufio"
	"io"
	"os"
)

import (
	"axaDB/src/dberrs"
  "axaDB/src/parsers"
)

func Connect(args []string) dberrs.AxaErr{

  if help := parsers.BoolParse(args, []string{"-h", "--help"}); help {
    fmt.Println(parsers.ConnectHelp())
    return dberrs.DB_NORM()
  }

  ip, err1  := parsers.InitParse(args, []string{"-i", "--ip"})
  if err1 != nil {
    ip = "localhost"
  }
  
  port, err2 := parsers.InitParse(args, []string{"-p", "--port"})
  if err2 != nil {
    port = "13131"
  }
  
  host, err := parsers.InitParse(args, []string{"-h", "--host"})
  if err != nil {
    host = ip + ":" + port
  } 

	conn, err := net.Dial("tcp", host)
  if err != nil {
    fmt.Println(dberrs.DB_C01(host).Err)
    return dberrs.DB_NORM()
  }

	var message string
	reader := bufio.NewReader(os.Stdin)
	wrapped_conn := bufio.NewReader(conn)
	buff := make([]byte, 256)

	for {
		fmt.Printf("axa exec $ ")
		message, _ = reader.ReadString('\n')
		message = message[:len(message)-1]
		m_len := byte(len(message))
		m_buff := make([]byte, 1)
		m_buff[0] = m_len

		_, err = conn.Write(append(m_buff[:], []byte(message)...))
		if err != nil {
			fmt.Println(err)
			return dberrs.DB_NORM()
		}	


		m_len, err := wrapped_conn.ReadByte()
    	if err != nil {
			conn.Close()
			break
    	}

		_, err = io.ReadFull(wrapped_conn, buff[:int(m_len)])
		if err != nil {
			conn.Close()
			return dberrs.DB_NORM()
		}

		fmt.Println(fmt.Sprintf("%s", buff[:int(m_len)]))

	}
	conn.Close()

	return dberrs.DB_NORM()
}
