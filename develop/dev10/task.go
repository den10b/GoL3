package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleClient(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции

	buf := make([]byte, 32) // буфер для чтения клиентских данных
	for {
		conn.Write([]byte("Hello, what's your name?\n")) // пишем в сокет

		readLen, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println(err)
			break
		}

		conn.Write(append([]byte("Goodbye, "), buf[:readLen]...)) // пишем в сокет
	}
}

func myTelnet(addr string, port int, timeout int) error {
	adres := addr + ":" + strconv.Itoa(port)
	fmt.Println(adres)
	//listener, err := net.Listen("tcp", adres) // открываем слушающий сокет
	dialTimeout, err := net.DialTimeout("tcp", adres, time.Duration(timeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer dialTimeout.Close()

	go func() {
		reader := bufio.NewReader(dialTimeout)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				fmt.Println(err)
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Print("Message from server: " + message)
		}
	}()

	//fmt.Println(listener)
	//beginTime := time.Now()
	//for time.Since(beginTime) < time.Duration(timeout)*time.Second {
	//	fmt.Println("Ща будем коннект")
	//	conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
	//	if err != nil {
	//		time.Sleep(time.Second)
	//		fmt.Println(err)
	//		continue
	//	}
	//	fmt.Println(conn)
	//	go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
	//}
	return nil
}

func main() {

	timeout := flag.Int("timeout", 10, "Таймаут")
	flag.Parse()
	var err error
	port, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	ip := os.Args[len(os.Args)-2]
	//port := 1232
	//ip := net.ParseIP("1.1.1.1")
	//timeout := 10
	err = myTelnet(ip, port, *timeout)
	if err != nil {
		log.Fatal(err)
	}

}
