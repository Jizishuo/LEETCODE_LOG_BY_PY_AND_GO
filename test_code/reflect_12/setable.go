package main

import (
	"fmt"
	"reflect"
)

func setable()  {
	x:=10
	v1 := reflect.ValueOf(x)
	fmt.Println("setable:", v1.CanSet())
	p:=reflect.ValueOf(&x)
	fmt.Println("setbale: ", p.CanSet())
	v2 := p.Elem()
	fmt.Println("setbale: ",v2.CanSet())
	// 结果
	// setable: false
	// setable: false
	// setable: true
}
func changetoseven(v reflect.Value)  {
	v.SetInt(7)
}

func cansetexample()  {
	x:=10
	v := reflect.ValueOf(&x).Elem()
	changetoseven(v)
	fmt.Println("value out side:", v.Interface())
	// 结果
	// value outside: 7
}

// 增加recover
func canNotSetExample() {
	x := 10
	v := reflect.ValueOf(x)
	changetoseven(v)
	fmt.Println("value outside:", v.Interface())

	// panic: reflect: reflect.flag.mustBeAssignable using unaddressable value

}

func main()  {
	// setable()
	// cansetexample()
	canNotSetExample()
}
