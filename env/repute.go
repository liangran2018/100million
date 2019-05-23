package env

var repute = 0

func ReputeAdd(i int) {
	repute += i
}

func ReputeSub(i int) {
	repute -= i
	if repute < 0 {
		repute = 0
	}
}

func ReputeGet() int {
	return repute
}
