package main

import (
	"net"
	"os"
	"os/signal"
	"runtime"
)

func main() {
	const CPUS = 1
	runtime.GOMAXPROCS(CPUS)

	listener, err := net.Listen("tcp", "localhost:10000")

	if err != nil {
		panic(err)
	}

	for i := 0; i < CPUS; i++ {
		go func() {
			for {
				con, err := listener.Accept()
				if err != nil {
					panic(err)
				}
				buf := make([]byte, 1024)
				con.Read(buf)
				go func() {
					defer con.Close()
				}()
			}
		}()
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
