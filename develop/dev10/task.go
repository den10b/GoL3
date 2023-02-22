package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func stopper(cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\ngot interrupt signal")
	cancelFunc()
	return
}
func reader(conn net.Conn, done context.Context, cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 32)

	for {
		select {
		default:
			_, err := conn.Read(buf)
			_, err = os.Stdout.Write(buf)

			if err != nil {
				cancelFunc()
				return
			}
			buf = make([]byte, 32)
		case <-done.Done(): //проверяем наличие чего-нибудбь в канале done
			return
		}

	}
}

func writer(conn net.Conn, done context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewReader(os.Stdin)
	for {
		select {
		default:
			t, err := scanner.ReadString('\n')
			if err != nil {
				return
			}
			_, err = conn.Write([]byte(t)) // пишем в сокет
			if err != nil {
				return
			}
		case <-done.Done(): //проверяем наличие чего-нибудбь в канале done
			return
		}

	}
}

func myTelnet(addr string, port int, timeout int) error {
	conn, err := net.DialTimeout("tcp", addr+":"+strconv.Itoa(port), time.Duration(timeout)*time.Second) // открываем слушающий сокет
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go reader(conn, ctx, cancel, wg)
	go stopper(cancel, wg)
	go writer(conn, ctx, wg)
	wg.Wait()
	return nil
}

func main() {
	var err error
	//go serv()
	timeout := flag.Int("timeout", 10, "Таймаут")
	flag.Parse()
	port, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	ip := os.Args[len(os.Args)-2]

	//port := 8083
	//ip := "127.0.0.1"
	//timeout := 10
	err = myTelnet(ip, port, *timeout)
	if err != nil {
		log.Fatal(err)
	}

}
