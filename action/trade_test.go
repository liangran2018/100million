package action

import (
	"testing"
	"github.com/liangran2018/100million/goods"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
)

func TestIsRepeat(t *testing.T) {
	ok := IsBuyRepeat(goods.Bag)
	t.Log(ok)

	ok = IsSellRepeat(goods.Bag)
	t.Log(ok)

	ok = IsBuyRepeat(goods.Bag)
	t.Log(ok)
}

func TestBuy(t *testing.T) {
	ok := Buy(goods.Bag, 100, 10)
	t.Log(ok)

	env.MoneyAdd(10000)
	ok = Buy(goods.Bag, 100, 10)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.Bag), own.GoodsPrice(goods.Bag))
	t.Log(env.MoneyGet(), env.HealthGet())

	ok = Buy(goods.Bag, 10, 10)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.Bag), own.GoodsPrice(goods.Bag))
	t.Log(env.MoneyGet(), env.HealthGet())

	s := Sell(goods.Bag, 200, 5)
	t.Log(s)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.Bag), own.GoodsPrice(goods.Bag))
	t.Log(env.MoneyGet(), env.HealthGet(), env.ReputeGet())

	s = Sell(goods.Bag, 300, 5)
	t.Log(s)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.Bag), own.GoodsPrice(goods.Bag))
	t.Log(env.MoneyGet(), env.HealthGet(), env.ReputeGet())
}
