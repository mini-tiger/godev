package main

type S struct {
	M *int //xxx 这里改为 值，则不会逃逸
	D map[int]int
}

func main() {
	var x S
	var i int
	var d = make(map[int]int) //xxx 逃逸到堆上
	ref(&i, &x, d)
}
func ref(y *int, z *S, d map[int]int) { //这里的z没有逃逸，xxx 而i却逃逸了，
	z.M = y
	z.D = d
}

//这是因为go的逃逸分析不知道z和i的关系，逃逸分析不知道参数y是z的一个成员，所以只能把它分配给堆。
