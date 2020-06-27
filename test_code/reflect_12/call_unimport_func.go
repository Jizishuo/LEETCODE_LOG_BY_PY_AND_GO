package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func (p *Person)getName() string  {
	return p.Name
}

func (p *Person)GetAge() int  {
	return p.Age
}

func main()  {
	p := &Person{
		"ji",
		20,
	}
	v := reflect.ValueOf(p)
	para := make([]reflect.Value, 0)
	age := v.MethodByName("GetAge").Call(para)
	fmt.Printf("age: %d\n", age[0].Interface().(int))

	name := v.MethodByName("getName").Call(para)
	fmt.Printf("name: %s\n", name[0].Interface().(string))
	// 要大写才能反射到
}