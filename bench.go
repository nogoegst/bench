// bench.go - benchmark readers and writers
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to bench, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package bench

import (
	"io"
	"io/ioutil"
	"math/rand"
	"time"
)

// Copy is a generic benchmarking function that copies data from r to w
// by chunks which sizes are specified in byteschedule.
// Copy returns all collected measures (even failed ones).
func Copy(w io.Writer, r io.Reader, byteschedule ...int64) (ms []Measure, err error) {
	for _, b := range byteschedule {
		start := time.Now()
		n, err := io.CopyN(w, r, b)
		elapsed := time.Since(start)
		m := Measure{
			Bytes:       n,
			Nanoseconds: elapsed.Nanoseconds(),
		}
		ms = append(ms, m)
		if err != nil {
			return ms, err
		}
	}
	return ms, err
}

// Read benchmarks io.Reader.
// Read reads data from r by chunks which sizes are specified in byteschedule.
// Read returns all collected measures (even failed ones).
func Read(r io.Reader, byteschedule ...int64) (ms []Measure, err error) {
	w := ioutil.Discard
	return Copy(w, r, byteschedule...)
}

// Write benchmarks io.Writer.
// Write writes random data to w by chunks which sizes are specified in byteschedule.
// Write returns all collected measures (even failed ones).
func Write(w io.Writer, byteschedule ...int64) (ms []Measure, err error) {
	// XXX: This may be slow. We need a random stream merely to prevent compression.
	r := rand.New(rand.NewSource(1).(rand.Source64))
	return Copy(w, r, byteschedule...)
}
