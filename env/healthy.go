package env

import "github.com/liangran2018/100million/base"

var health = 100

func HealthAdd(i int) {
	health += i
	if health > 100 {
		health = 100
	}
}

func HealthSub(i int) {
	health -= i
	if health <= 0 {
		panic(health)
	}

	if health < 50 && base.Rand(100) >= health {
		Ill()
	}
}

func HealthWhole() {
	health = 100
}

func Ill() {}

func HealthGet() int {
	return health
}