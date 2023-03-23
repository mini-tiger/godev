package main

import (
	"fmt"
	// "sync"
	"io"
	"os"
	"reflect"
	// "runtime"
	// "time"
)

func main() {

	t := reflect.TypeOf(3)
	fmt.Println(t, t.String(), t.Size(), t.Kind()) //int int 8 int

	var (
		w io.Writer = os.Stdout
	)
	ww := reflect.TypeOf(w)
	fmt.Println(ww)
	fmt.Printf("%T\n", os.Stdout) //*os.File  %T 是reflect.TypeOf简写

	type Key struct {
		Path, Country string
	}
	m := map[Key]int{
		Key{"1", "2"}: 11,
		Key{"2", "3"}: 12,
	}
	mm := reflect.TypeOf(m)
	fmt.Println(mm, mm.String(), mm.Kind(), mm.Elem()) //map[main.Key]int map[main.Key]int map

	v := reflect.ValueOf(m)
	fmt.Println(v)

	// for _, key := range v.MapIndex() {
	// 	fmt.Println(key, key.String(), key.Kind())
	// }

	m1 := map[string]int32{
		"a": 1,
		"b": 2,
	}
	mm1 := reflect.ValueOf(m1)
	for i, vv := range mm1.MapKeys() {
		fmt.Println(i, vv.String())
	}
	fmt.Println("=========")
	mm1Range := mm1.MapRange()
	for {
		if !mm1Range.Next() {
			break

		}
		fmt.Println(mm1Range.Key().String(), mm1Range.Value())
	}

}
