package env

var money = 0

func MoneyGet() int {
	return money
}

func MoneyAdd(i int) {
	money += i
}

func MoneySub(i int) {
	money -= i
	if money < 0 {
		money = 0
	}
}