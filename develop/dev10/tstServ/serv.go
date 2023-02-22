package main

import (
	"log"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8083") // открываем слушающий сокет
	for {
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		if err != nil {
			continue
		}
		go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции

	buf := make([]byte, 32) // буфер для чтения клиентских данных
	for {
		conn.Write([]byte("SERVER IS READY!!!\n")) // пишем в сокет
		readLen, err := conn.Read(buf)             // читаем из сокета
		if err != nil {
			log.Fatalln(err)
		}

		resBytes := append([]byte("SERVER RECIEVED A MESSAGE: "), buf[:readLen]...)
		resBytes = append(resBytes, []byte("\n")...)
		conn.Write(resBytes) // пишем в сокет
		buf = make([]byte, 32)
	}
}
