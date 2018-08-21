package g

type Food interface {
	zhengqian() float64
}

type Cook_food struct {
	Food map[string]Food

}

type Milk struct {
	RMB float64
	Chengben float64
}

type Fish struct {
	Name string
	RMB float64
	Chengben float64
}

