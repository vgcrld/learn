package employee

type Boss struct {
	name   string
	title  string
	salary float32
}

func Salary(base, raise float32) float32 {
	return base * raise
}
