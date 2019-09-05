package main

import (
	"net"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	listener, err := net.Listen("tcp", "localhost:10000")

	if err != nil {
		panic(err)
	}

	go func() {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		buf := make([]byte, 1024)
		con.Read(buf)
		go func() {
			defer con.Close()
		}()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
