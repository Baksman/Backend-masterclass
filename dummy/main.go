package main

import (
	"fmt"
	"reflect"
)

type Stringer interface {
	namer(name string) string
}
type newInt string

func (a newInt) namer(name string) string {
	return name
}
func main() {
	var a Stringer = newInt("hello")
	type A = [16]int16
	var c <-chan map[A][]byte
	_ = c
	tc := reflect.TypeOf(a)
	fmt.Println(tc.Kind())
	fmt.Println(tc)
	fmt.Println(tc.ChanDir())

	// var a Stringer = newInt(10)

}
