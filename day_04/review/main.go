package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Value of %v is %v, typeName is %v, typeKind is %v\n", x, v, v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Printf("%T, %v\n", k, k)
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
		if(k == reflect.Float32){
			fmt.Println("Yes reflect")
		}

	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

type student struct{}
type myInt int16

func main() {
	var a float32 = 3.14
	// reflectType(a)
	reflectValue(a)
	var b myInt = 6
	// reflectType(b)
	reflectValue(b)
	c := student{}
	// reflectType(c)
	reflectValue(c)
}
