package today

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

func TodayCompany() {
	for k := range companySave {
		p := base.Rand(100)
		companySave[k].Price, companySave[k].IsSt = company.GetPriceByNews(companySave[k].Company, p, companySave[k].Price)
	}
}

func GetTodayCompany() []*NowCompany {
	return companySave
}

func TodayCompanyShow() {
	for k, v := range companySave {
		if v.IsSt {
			base.Notice(fmt.Sprintf("%d. %s: 退市", k+1, v.Company.Name()))
			continue
		}
		base.Notice(fmt.Sprintf("%d. %s 价格: %.2f", k+1, v.Company.Name(), v.Price))
	}
}

