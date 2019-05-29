package action

import (
	"github.com/liangran2018/100million/company"
	"github.com/liangran2018/100million/env"
	"github.com/liangran2018/100million/own"
)

var companyact = make(map[int]map[company.CompanyIndex]*act)

func Invest(c company.CompanyIndex, price, num float32) string {
	if env.MoneyGet() < int(price * num) {
		return "money not enough"
	}

	own.Invest(c, price, num)
	env.MoneySub(int(price * num))

	if !IsInvestRepeat(c) {
		env.HealthSub(1)
		env.ReputeAdd(5)
	}

	return "ok"
}

func DisInvest(c company.CompanyIndex, price, num float32) string {
	if own.CompanyNum(c) < int(num) {
		return "num wrong"
	}

	own.DisInvest(c, price, num)
	env.MoneyAdd(int(price * num))

	if !IsDisInvestRepeat(c) {
		env.HealthSub(1)
	}
	return "ok"
}

func IsInvestRepeat(c company.CompanyIndex) bool {
	age := env.GetAge()
	gs, ok := companyact[age]
	if !ok {
		companyact[age] = make(map[company.CompanyIndex]*act)
		companyact[age][c] = &act{buy:true}
		return false
	}

	a, ok := gs[c]
	if !ok {
		companyact[age][c].buy = true
		return false
	}

	if a.buy {
		return true
	}

	companyact[age][c].buy = true
	return false
}

func IsDisInvestRepeat(c company.CompanyIndex) bool {
	age := env.GetAge()
	gs, ok := companyact[age]
	if !ok {
		companyact[age] = make(map[company.CompanyIndex]*act)
		companyact[age][c] = &act{sell:true}
		return false
	}

	a, ok := gs[c]
	if !ok {
		companyact[age][c].sell = true
		return false
	}

	if a.sell {
		return true
	}

	companyact[age][c].sell = true
	return false
}