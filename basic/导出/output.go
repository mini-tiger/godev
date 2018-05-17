package 导出
const (
	a = 1  //iota 0, 小写不能导出
	B	//1  大写可以导出,没有定义数值和上面a一样是1
	C	//2
	D = 'A'
	_E = "123" //常量名下划线开头不能导出
	f = len(_E)//必须是已内建函数
	g = iota//iota 无论何时定义，都是从const本组内 0至当前的行位置 ，当前行位置6
)

const   (
	aa,BB,CC=2,"b","c"
	dd,EE,FF//同时定义3个，与上面数量一样，和上面一样 2 b c
	GG=iota//当前行位置 2

)

var N  ="nnn"
