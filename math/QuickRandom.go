package math

import "time"

type QuickRandom uint64

func (r QuickRandom) GetN(n uint64) uint64 {
	return uint64(r) % n
}

func (r QuickRandom) Next() QuickRandom {
	return r*214013 + 2531011
}

var rnd = QuickRandom(time.Now().UnixNano())

func RandomU32() uint32 {
	rnd = rnd.Next()
	return uint32(rnd >> 8)
}
