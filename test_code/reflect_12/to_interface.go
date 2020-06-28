package main

import (
	"fmt"
	"reflect"
)

func printvv(v reflect.Value)  {
	fmt.Println("value: ", v.Interface())
}

func main()  {
	a :=1
	b := 3.1415926
	c := "tree"
	d := false
	e := []int{1,2,3,4,5}
	printvv(reflect.ValueOf(a))
	printvv(reflect.ValueOf(b))
	printvv(reflect.ValueOf(c))
	printvv(reflect.ValueOf(d))
	printvv(reflect.ValueOf(e))

	type myint int
	var x myint = 10
	printvv(reflect.ValueOf(x))
	fmt.Println("type: ", reflect.TypeOf(x))

	// value:  1
	//value:  3.1415926
	//value:  tree
	//value:  false
	//value:  [1 2 3 4 5]
	//value:  10
	//type:  main.myint
}