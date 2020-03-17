package main

type S struct {
	M *int // xxx 这里换成值，可以防止逃逸
}

func main() {
	var i int
	var ss *S
	ss.refStruct(i)
}
func (s *S) refStruct(y int) {
	s.M = &y
}

//这里的y是逃逸了，与是否在struct里好像并没有区别，有可能被函数外的程序访问就会逃逸
