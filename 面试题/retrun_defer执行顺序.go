package main

/**
 * @Author: Tao Jun
 * @Description: main
 * @File:  retrun_defer执行顺序
 * @Version: 1.0.0
 * @Date: 2021/4/28 上午10:03
 */

func main() {
	println(f1())
	println(f2())
	println(f3())
}

/*
1. t 先定义就是t=0
2. return    t 这时是0
3. defer  t++
4. t 是1
*/

func f1() (t int) { // 返回1,  调试模式
	defer func() {
		t++
	}()
	return t
}

/*
1. 没有先定义   临时变量是0
2. return    t 这时是5, 临时变量=t , 临时变量是5
3. defer  t++    ;不影响临时变量
4. 临时变量是5
*/

func f2() int { // 返回1,  调试模式
	t := 5
	defer func() {
		t++
	}()
	return t
}

/*
defer r 传入参数 是局部变量

*/
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
