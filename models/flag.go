package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") //传递内存地址
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(n, *n)
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*cmd
C:\godev\models>go run flag.go -s // -n a b c
[a b c]
0xc04200823a true
a//b//c

*/
