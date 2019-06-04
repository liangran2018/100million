package company

import "github.com/liangran2018/100million/base"

type companyFeature struct {
	name     string
	born, st float32
	middle   int
	good     []oneNew
	bad      []oneNew
}

type oneNew struct {
	p base.Pro
	c float32
}

type CompanyIndex int

var companys map[CompanyIndex]companyFeature

const (
	Longtu CompanyIndex = iota
	CompanysEnd
)

func init() {
	companys = make(map[CompanyIndex]companyFeature, CompanysEnd)
	companys[Longtu] = companyFeature{}
}

func GetPriceByNews(c CompanyIndex, pro int, price float32) (float32, bool) {
	cf := companys[c]

	if cf.middle <= pro {
		if cf.good[1].p.Min > pro {
			price = price*1.1 + cf.good[0].c
			return price, price <= cf.st
		} else if cf.good[1].p.Max < pro {
			price = price*2 + cf.good[2].c
			return price, price <= cf.st
		} else {
			price = price*1.5 + cf.good[1].c
			return price, price <= cf.st
		}
	} else {
		if cf.bad[1].p.Min > pro {
			price = price*0.9 - cf.bad[0].c
			return price, price <= cf.st
		} else if cf.bad[1].p.Max < pro {
			price = price/2 - cf.bad[2].c
			return price, price <= cf.st
		} else {
			price = price*0.7 - cf.bad[1].c
			return price, price <= cf.st
		}
	}
}

func (g CompanyIndex) Name() string {
	return companys[g].name
}

func (g CompanyIndex) Born() float32 {
	return companys[g].born
}

func (g CompanyIndex) St() float32 {
	return companys[g].st
}
