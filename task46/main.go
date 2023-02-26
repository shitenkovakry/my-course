package main

import (
	"fmt"
	"net"
)

const (
	stopWord = "stop"
)

func main() {
	// kotik podkliuchaet stakan k portu 1234 i derzhit stakan v peremennoi "listener"
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// kotik prilozhil stakan k uhu i zhdet podklichenie
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	for {
		// slushaem byte iz podkliuchenia i stavim v peremennoi "read"
		read := make([]byte, 64)
		n, err := conn.Read(read)
		if err != nil {
			panic(err)
		}

		fmt.Println(n, string(read))

		potentialStopWord := read[0:len(stopWord)]
		if string(potentialStopWord) == stopWord {
			break
		}

		nWrote, err := conn.Write([]byte(fmt.Sprintf("myaf:%s", string(read))))
		if err != nil {
			panic(err)
		}

		fmt.Println("wrote:", nWrote)
	}

	nWrote, err := conn.Write([]byte("bye!"))
	if err != nil {
		panic(err)
	}

	fmt.Println("wrote:", nWrote)

	// priniali vce i polozhili stakan (trubku)
	if err := conn.Close(); err != nil {
		panic(err)
	}
}
