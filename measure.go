// measure.go - results of bandwith measures.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to bench, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package bench

import (
	"time"
)

// Measure signifies result of a bandwidth measure.
// There were Bytes bytes sent/received for Nanoseconds nanoseconds.
type Measure struct {
	Bytes       int64
	Nanoseconds int64
}

func (m Measure) Per(t time.Duration) int64 {
	return m.Bytes * int64(t) / m.Nanoseconds
}

func (m Measure) PerHour() int64 {
	return m.Per(time.Hour)
}

func (m Measure) PerMinute() int64 {
	return m.Per(time.Minute)
}

func (m Measure) PerSecond() int64 {
	return m.Per(time.Second)
}

func (m Measure) PerMillisecond() int64 {
	return m.Per(time.Millisecond)
}

func (m Measure) PerMicrosecond() int64 {
	return m.Per(time.Microsecond)
}

func (m Measure) PerNanosecond() int64 {
	return m.Per(time.Nanosecond)
}
