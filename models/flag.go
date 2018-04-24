package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

//http://blog.studygolang.com/2013/02/%E6%A0%87%E5%87%86%E5%BA%93-%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%8F%82%E6%95%B0%E8%A7%A3%E6%9E%90flag/
var n = flag.Bool("n", false, "omit trailing newline") //传递内存地址
var sep = flag.String("s", " ", "separator")
var period = flag.Duration("period", 10*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Println(flag.Args()) //[a b c]
	fmt.Println(n, *n)
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Println()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
}

/*cmd
[a b c]
0xc04200c22a true
a//b//c
Sleeping for 10s...

*/
