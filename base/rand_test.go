package base

import "testing"

func TestRand(t *testing.T) {
	for i := 0; i < 100; i++ {
		j := Rand(1000)
		t.Log(j)
	}
}

func TestRandMany(t *testing.T) {
	i := RandMany(6, 5)
	t.Log(i)
}
