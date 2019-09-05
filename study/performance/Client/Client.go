package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(6)
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			var con net.Conn
			//var err error
			for {
				con, _ = net.Dial("tcp", "localhost:10000")

				if con != nil {
					break
				}
				time.Sleep(10 * time.Microsecond)
			}
			defer func() {
				con.Close()
				wg.Done()
			}()
			_, _ = con.Write([]byte("Hello"))
		}()
	}
	wg.Wait()
	fmt.Printf("%dNs\n", (time.Now().Sub(start)).Nanoseconds())
}
