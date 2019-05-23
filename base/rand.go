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

func RandMany(total, need int) []int {
	if total <= need {
		g := make([]int, total)
		for i := 0; i < total; i++ {
			g[i] = i
		}
		return g
	}

	g := make([]int, 0)

	for {
		i := Rand(total)
		if !In(i, g) {
			g = append(g, i)
		}

		if len(g) == need {
			break
		}
	}

	return g
}
