package action

import (
	"testing"
	"github.com/liangran2018/100million/goods"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
)

func TestIsRepeat(t *testing.T) {
	ok := IsRepeat(goods.MilkTee, "B")
	t.Log(ok)

	ok = IsRepeat(goods.MilkTee, "S")
	t.Log(ok)

	ok = IsRepeat(goods.MilkTee, "B")
	t.Log(ok)
}

func TestBuy(t *testing.T) {
	ok := Buy(goods.MilkTee, 100, 10)
	t.Log(ok)

	env.MoneyAdd(10000)
	ok = Buy(goods.MilkTee, 100, 10)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.MilkTee), own.GoodsPrice(goods.MilkTee))
	t.Log(env.MoneyGet(), env.HealthGet())

	ok = Buy(goods.MilkTee, 10, 10)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.MilkTee), own.GoodsPrice(goods.MilkTee))
	t.Log(env.MoneyGet(), env.HealthGet())

	ok = Sell(goods.MilkTee, 200, 5)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.MilkTee), own.GoodsPrice(goods.MilkTee))
	t.Log(env.MoneyGet(), env.HealthGet(), env.ReputeGet())

	ok = Sell(goods.MilkTee, 300, 5)
	t.Log(ok)
	t.Log(own.RoomUsed(), own.RoomFree())
	t.Log(own.GoodsNum(goods.MilkTee), own.GoodsPrice(goods.MilkTee))
	t.Log(env.MoneyGet(), env.HealthGet(), env.ReputeGet())
}
