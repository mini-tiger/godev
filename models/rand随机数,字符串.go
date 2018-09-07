package main

import (
	"fmt"
	"math/rand"
	zrand "crypto/rand"
	"bytes"
	"math/big"
	"strings"
)

const (
	only_num   = iota
	only_lower
	only_upper
	all
	initnum    = "0123456789"
	initlower  = "abcdefghijklmnopqrstuvwxyz"
)

var initupper string = strings.ToUpper(initlower)

func randstr(n, t int) (string, error) {
	if t > 3 {
		return "", fmt.Errorf("未定义的类型")
	}
	bs := make([]string, n) //初始化随机字符串的 切片
	rand.Seed(100)          //默认资源初始化状态, 种子给定
	c_rand_strfunc := func(is []string) {
		for i := 0; i < n; i++ {
			bs[i] = is[rand.Intn(len(is)-1)] //随机切片每个元素 从传入切片内 随机填充
		}
	}
	switch {
	case t == 0: // todo 只有数字
		in := strings.Split(initnum, "")
		c_rand_strfunc(in)
	case t == 1:
		is := strings.Split(initlower, "")
		c_rand_strfunc(is)
	case t == 2:
		is := strings.Split(initupper, "")
		c_rand_strfunc(is)
	case t == 3:
		in := strings.Split(initnum, "")
		is := strings.Split(initlower, "")
		is1 := strings.Split(initupper, "")
		in = append(in, is...)
		in = append(in, is1...)
		c_rand_strfunc(in)
	}
	fmt.Println(strings.Join(bs, ""))
	return strings.Join(bs, ""), nil
}
func main() {
	jia()  //伪随机数
	zhen() //真随机数
	fmt.Println("=================随机字符串====================")
	randstr(10, only_num) //10w位,只有数字的随机字符串
	randstr(10, only_lower)
	randstr(10, only_upper)
	randstr(10, all)
}

func zhen() {
	fmt.Println("=================真随机数====================")
	b := bytes.NewBuffer([]byte{1, 2})
	bint := big.NewInt(1000)
	n, err := zrand.Int(b, bint)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Println(n)
}
func jia() {
	fmt.Println(rand.Int())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Uint32())
	fmt.Println(rand.Int31n(1000)) //给定 0 - n 取值范围
	fmt.Println(rand.Float32())
	fmt.Println(rand.Perm(5)) // n 个元素的随机数切片

	//todo 有确定性生成随机数，每次生成会是一样的
	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	fmt.Println(r1.Int()) //todo 与下面r2生成的一样，7530908113823513298

	s2 := rand.NewSource(100)
	r2 := rand.New(s2)
	fmt.Println(r2.Int())
}
