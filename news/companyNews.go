package news

import (
	"github.com/liangran2018/100million/company"
	"github.com/liangran2018/100million/base"
	"github.com/liangran2018/100million/goods"
)

type NowCompany struct {
	Company company.CompanyIndex
	News companysNewIndex
	Price float32
}

type companysNewIndex int

var companyNews map[company.CompanyIndex][]companysNewIndex
var companyNewsStr map[companysNewIndex]newsFeature
var companySave []*NowCompany

const (
	CompanyAll companysNewIndex = iota
	LongtuGood
	LongtuBad
	companyEnd
)

func init() {
	companyNews = make(map[company.CompanyIndex][]companysNewIndex, company.CompanysEnd)
	companyNews[company.Longtu] = []companysNewIndex{CompanyAll, LongtuGood, LongtuBad}

	companyNewsStr = make(map[companysNewIndex]newsFeature, companyEnd)

	companySave = make([]*NowCompany, 0, len(companyNews))
	for k, _ := range companyNews {
		save := &NowCompany{Company:k, Price:k.Born()}
		companySave = append(companySave, save)
	}
}

func (n companysNewIndex) Intro() string {
	return companyNewsStr[n].intro
}

func (n companysNewIndex) Style() int {
	return companyNewsStr[n].style
}

func (n companysNewIndex) pro() (max, min int) {
	p := companyNewsStr[n]
	return p.promax, p.promin
}

func companyGet() {
	for k, gIdx := range companyNews {
		save := &NowCompany{Company:k}
		p := base.Rand(100)
		for _, nIdx := range companyNews[k] {
			pmax, pmin := nIdx.pro()
			if p >= pmin && p < pmax {
				save.News = nIdx
				switch nIdx.Style() {
				case styleGood:
					save.Price += save.Price * 0.1 + save.Company.Up()[0]
				case styleBetter:
					save.Price += save.Price * 0.5 + save.Company.Up()[1]
				case styleBest:
					save.Price += save.Price + save.Company.Up()[2]
				case styleBad:
					save.Price = save.Price * 0.1
				}
				if nIdx.Style() == styleGood {

				} else if nIdx.Style() == styleBad {
					save[k].Price = save[k].Goods.Min()
				} else {
					save[k].Price = save[k].Goods.Normal()
				}
			}
		}
	}
}

func GetMarket() []*NowMarket {
	return save
}

func MarketNewsShow() []string {
	marketGet()

	ns := make([]newsIndex, 0)
	for _, v := range save {
		ns = append(ns, v.News)
	}

	news := make([]string, 0)
	for k := range news {
		if ns[k].Style() == styleNormal {
			continue
		} else {
			news = append(news, ns[k].Intro())
		}
	}

	if len(news) == 0 {
		news = append(news, newsStr[All].intro)
	}

	return news
}

