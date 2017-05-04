// zeronull.go - benchmark /dev/zero and /dev/null
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to bench, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"log"
	"os"

	"github.com/nogoegst/bench"
)

func main() {
	r, err := os.Open("/dev/zero")
	if err != nil {
		log.Fatal(err)
	}
	ms, err := bench.Read(r, 2, 432432432, 3243243, 234324344, 17)
	if err != nil {
		log.Print(err)
	}
	for _, m := range ms {
		log.Printf("%v bytes/sec", m.PerSecond())
	}

	w, err := os.OpenFile("/dev/null", os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	ms, err = bench.Write(w, 2, 432432432, 3243243, 234324344, 17)
	if err != nil {
		log.Print(err)
	}
	for _, m := range ms {
		log.Printf("%v bytes/sec", m.PerSecond())
	}
}
