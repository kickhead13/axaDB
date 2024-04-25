package axa_connect

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

import (
	"axaDB/src/dberrs"
)

func Connect(args []string) dberrs.AxaErr{


	conn, err := net.Dial("tcp", "localhost:13131")
    if err != nil {
        fmt.Println(err)
        return dberrs.DB_NORM()
    }

	var message string
	reader := bufio.NewReader(os.Stdin)

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
	}
	conn.Close()

	return dberrs.DB_NORM()
}