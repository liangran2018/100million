package base

import (
	"time"
	"math/rand"
)

var rd = rand.NewSource(0)

func Rand(tt int) int {
	rd.Seed(time.Now().UnixNano())
	r := rand.New(rd)
	return r.Intn(tt)
}

func MMrand(max, min int) int {
	return Rand(max-min) + min
}
