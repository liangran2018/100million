package own

import (
	"github.com/liangran2018/100million/goods"
)

type ownsFeature struct {
	num int
	price int
}

var room = 100
var own map[goods.GoodsIndex]*ownsFeature

func RoomGet() int {
	return room
}

func RoomAdd(i int) {
	room += i
}

func RoomUsed() int {
	r := 0
	for _, v := range own {
		r += v.num
	}

	return r
}

func RoomFree() int {
	return room - RoomUsed()
}

func init() {
	own = make(map[goods.GoodsIndex]*ownsFeature)
}

func Buy(g goods.GoodsIndex, price, num int) {
	f, ok := own[g]
	if !ok {
		own[g] = &ownsFeature{price:price, num:num}
	} else {
		f.price = (f.price * f.num + price * num) / (f.num + num)
		f.num += num
	}
}

func Sell(g goods.GoodsIndex, num int) {
	own[g].num -= num
	if own[g].num == 0 {
		delete(own, g)
	}
}

func GoodsPrice(g goods.GoodsIndex) int {
	f, ok := own[g]
	if !ok {
		return 0
	}
	return f.price
}

func GoodsNum(g goods.GoodsIndex) int {
	f, ok := own[g]
	if !ok {
		return 0
	}
	return f.num
}