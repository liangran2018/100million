package own

import (
	"github.com/liangran2018/100million/company"
)

type NowInvest struct {
	Num float32
	Price float32
}

var owncompany = make(map[company.CompanyIndex]*NowInvest)

func GetOwnCompany() map[company.CompanyIndex]*NowInvest {
	return owncompany
}

func Invest(c company.CompanyIndex, price, num float32)  {
	f, ok := owncompany[c]
	if !ok {
		owncompany[c] = &NowInvest{Price:price, Num:num}
	} else {
		f.Price = (f.Price * f.Num + price * num) / (f.Num + num)
		f.Num += num
	}
}

func DisInvest(c company.CompanyIndex, price, num float32) {
	if owncompany[c].Num == num {
		delete(owncompany, c)
		return
	}

	owncompany[c].Price = (owncompany[c].Price * owncompany[c].Num - price * num) / (owncompany[c].Num - num)
	owncompany[c].Num -= num
}

func St(c company.CompanyIndex) int {
	p := owncompany[c].Num * c.St() / 2
	delete(owncompany, c)
	return int(p)
}

func CompanyTotalPrice() int {
	var t float32
	for _, v := range owncompany {
		t += v.Price * v.Num
	}

	return int(t)
}

func CompanyNum(c company.CompanyIndex) int {
	f, ok := owncompany[c]
	if !ok {
		return 0
	}
	return int(f.Num)
}