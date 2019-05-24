package own

import (
	"github.com/liangran2018/100million/goods"
)

type NowStore struct {
	Goods goods.GoodsIndex
	Num int
	Price int
}

var room = 100
var own map[goods.GoodsIndex]*NowStore

func RoomGet() int {
	return room
}

func RoomAdd(i int) {
	room += i
}

func RoomUsed() int {
	r := 0
	for _, v := range own {
		r += v.Num
	}

	return r
}

func RoomFree() int {
	return room - RoomUsed()
}

func init() {
	own = make(map[goods.GoodsIndex]*NowStore)
}

func Buy(g goods.GoodsIndex, price, num int) {
	f, ok := own[g]
	if !ok {
		own[g] = &NowStore{Goods:g, Price:price, Num:num}
	} else {
		f.Price = (f.Price * f.Num + price * num) / (f.Num + num)
		f.Num += num
	}
}

func Sell(g goods.GoodsIndex, num int) {
	own[g].Num -= num
	if own[g].Num == 0 {
		delete(own, g)
	}
}

func GoodsPrice(g goods.GoodsIndex) int {
	f, ok := own[g]
	if !ok {
		return 0
	}
	return f.Price
}

func GoodsNum(g goods.GoodsIndex) int {
	f, ok := own[g]
	if !ok {
		return 0
	}
	return f.Num
}

func GoodsTotalPrice() int {
	t := 0
	for _, v := range own {
		t += v.Price * v.Num
	}

	return t
}

func GetStore() []*NowStore {
	s := make([]*NowStore, 0, len(own))
	for _, v := range own {
		s = append(s, v)
	}

	return s
}
