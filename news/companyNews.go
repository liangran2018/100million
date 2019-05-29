package news

import (
	"github.com/liangran2018/100million/company"
	"github.com/liangran2018/100million/base"
	"fmt"
)

type NowCompany struct {
	Company company.CompanyIndex
	Price float32
	IsSt  bool
}

var companySave []*NowCompany

func init() {
	companySave = make([]*NowCompany, company.CompanysEnd)
	for i := company.Longtu; i < company.CompanysEnd; i++ {
		companySave[i] = &NowCompany{Company:i, Price:i.Born()}
	}
}

func NewCompany() {
	for k := range companySave {
		p := base.Rand(100)
		companySave[k].Price, companySave[k].IsSt = company.GetPriceByNews(companySave[k].Company, p, companySave[k].Price)
	}
}

func GetCompany() []*NowCompany {
	return companySave
}

func CompanyShow() []string {
	s := make([]string, len(companySave))
	for k := range companySave {
		s[k] = fmt.Sprintf("%s %d ", companySave[k].Company.Name(), companySave[k].Price)
	}


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

