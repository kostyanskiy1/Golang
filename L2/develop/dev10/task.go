package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type Cmd struct {
	Timeout uint
}

func (cmd *Cmd) Usage() {
	fmt.Printf("Usage of %s:\ngo run task.go [ -timeout [timeout]] [host] [port] \n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.UintVar(&cmd.Timeout, "timeout", 10, "максимальное время ожидания сервера")
	flag.Parse()
}

func main() {
	var conn net.Conn
	var err error

	stop := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	startTime := time.Now()
	var cmd = Cmd{}
	cmd.Parse()

	if len(flag.Args()) < 2 {
		log.Fatalf("Need enter host and port!\nUse command in format: telnet [-timeout [timeout]] [host] [port]")
		return
	}

	go func() {
		<-stop
		log.Fatalf("Сигнал для остановки принят. Закрытие сокета и завершение программы...")
		// Код для закрытия сокета, остановки сервера или другого необходимого завершения
		done <- true

	}()

	// Connect to socket (default timeout = 10s)
	for {
		conn, err = net.Dial("tcp", flag.Args()[0]+":"+flag.Args()[1]) //socket
		if err == nil {
			log.Printf("Connected to server %v:%v", flag.Args()[0], flag.Args()[1])
			break
		} else if time.Since(startTime) > time.Duration(cmd.Timeout)*time.Second {
			log.Fatal("Connection failed: timeout expired.")
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		//copyTo(os.Stdout, conn)
		if err != nil {
			break
		}
		fmt.Print("Message from server: " + message)
	}
	<-done
	defer conn.Close()
}
