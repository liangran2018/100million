package action

import (
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/goods"
	"github.com/liangran2018/100million/own"
)

type act struct {
	buy bool
	sell bool
}

var goodsact = make(map[int]map[goods.GoodsIndex]*act)

func Buy(goods goods.GoodsIndex, price, num int) string {
	if env.MoneyGet() < price * num {
		return "money not enough"
	}

	if num > own.RoomFree() {
		return "room not enough"
	}

	own.Buy(goods, price, num)
	env.MoneySub(price * num)

	if !IsBuyRepeat(goods) {
		env.HealthSub(1)
	}

	return "ok"
}

func Sell(goods goods.GoodsIndex, price, num int) string {
	if own.GoodsNum(goods) < num {
		return "num wrong"
	}

	own.Sell(goods, num)
	env.MoneyAdd(price * num)

	if !IsSellRepeat(goods) {
		env.HealthSub(1)
		if own.GoodsPrice(goods) < price {
			env.ReputeAdd(2)
		}
	}
	return "ok"
}

func IsBuyRepeat(g goods.GoodsIndex) bool {
	age := env.GetAge()
	gs, ok := goodsact[age]
	if !ok {
		goodsact[age] = make(map[goods.GoodsIndex]*act)
		goodsact[age][g] = &act{buy:true}
		return false
	}

	a, ok := gs[g]
	if !ok {
		goodsact[age][g].buy = true
		return false
	}

	if a.buy {
		return true
	}

	goodsact[age][g].buy = true
	return false
}

func IsSellRepeat(g goods.GoodsIndex) bool {
	age := env.GetAge()
	gs, ok := goodsact[age]
	if !ok {
		goodsact[age] = make(map[goods.GoodsIndex]*act)
		goodsact[age][g] = &act{sell:true}
		return false
	}

	a, ok := gs[g]
	if !ok {
		goodsact[age][g].sell = true
		return false
	}

	if a.sell {
		return true
	}

	goodsact[age][g].sell = true
	return false
}