// inet.go - benchmark TCP connections.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to bench, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"log"
	"net"

	"github.com/nogoegst/bench"
)

func main() {
	host := "speedtest.local:9999"
	bs := []int64{24, 5435, 3434, 555533334}

	func() {
		c, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		ms, err := bench.Read(c, bs...)
		if err != nil {
			log.Print(err)
		}
		log.Printf("read:")
		for _, m := range ms {
			log.Printf("%v bytes/sec", m.PerSecond())
		}
	}()
	func() {
		c, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		ms, err := bench.Write(c, bs...)
		if err != nil {
			log.Print(err)
		}
		log.Printf("write:")
		for _, m := range ms {
			log.Printf("%v bytes/sec", m.PerSecond())
		}
	}()
}
