package company

type companyFeature struct {
	name      string
	intro     string
	born      float32
	up1       float32
	up2       float32
	up3       float32
	down1     float32
	down2     float32
	down3     float32
	st        float32
}

type CompanyIndex int

var companys map[CompanyIndex]companyFeature

const (
	Longtu = iota
	CompanysEnd
)

func init() {
	companys = make(map[CompanyIndex]companyFeature, CompanysEnd)
	companys[Longtu] = companyFeature{}
}

func (g CompanyIndex) Name() string {
	return companys[g].name
}

func (g CompanyIndex) Intro() string {
	return companys[g].intro
}

func (g CompanyIndex) Born() float32 {
	return companys[g].born
}

func (g CompanyIndex) Up() []float32 {
	c := companys[g]
	return []float32{c.up1, c.up2, c.up3}
}

func (g CompanyIndex) Down() []float32 {
	c := companys[g]
	return []float32{c.down1, c.down2, c.down3}
}

func (g CompanyIndex) St() float32 {
	return companys[g].st
}
