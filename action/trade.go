package action

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/goods"
	"github.com/liangran2018/100million/own"
)

var act map[int]map[goods.GoodsIndex][]string

func init() {
	act = make(map[int]map[goods.GoodsIndex][]string)
}

func Buy(goods goods.GoodsIndex, price, num int) bool {
	if env.MoneyGet() < price * num {
		return false
	}

	if num > own.RoomFree() {
		return false
	}

	own.Buy(goods, price, num)
	env.MoneySub(price * num)

	if !IsRepeat(goods, "B") {
		env.HealthSub(1)
	}

	return true
}

func Sell(goods goods.GoodsIndex, price, num int) bool {
	if own.GoodsNum(goods) < num {
		return false
	}

	own.Sell(goods, num)
	env.MoneyAdd(price * num)

	if !IsRepeat(goods, "S") {
		env.HealthSub(1)
		if own.GoodsPrice(goods) < price {
			env.ReputeAdd(2)
		}
	}
	return true
}

func IsRepeat(g goods.GoodsIndex, s string) bool {
	age := env.GetAge()
	gs, ok := act[age]
	if !ok {
		act[age] = make(map[goods.GoodsIndex][]string)
		act[age][g] = []string{s}
		return false
	}

	a, ok := gs[g]
	if !ok {
		act[age][g] = []string{s}
		return false
	}

	for _, v := range a {
		if v == s {
			return true
		}
	}

	act[age][g] = append(act[age][g], s)
	return false
}
