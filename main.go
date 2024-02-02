package main

import (
	"fmt"
	"reflect"
)

func merge(arr interface{}, nn interface{}) (interface{},error) {

	if arr == nil {
		if nn == nil {
		return nil, fmt.Errorf("Nil Input")
		} else {
			nnReflected := reflect.ValueOf(nn)

			switch nnReflected.Kind() {
			case reflect.Slice:
				return nn,nil
			default :
				return []interface{}{nn},nil
			}
		}
	}

	var out []interface{}

	arrReflected := reflect.ValueOf(arr)
	nnReflected := reflect.ValueOf(nn)

	switch arrReflected.Kind() {
	case reflect.Slice :
		for i:= 0; i < arrReflected.Len(); i++ {
			out = append(out,arrReflected.Index(i).Interface())
		}
	default:
		out = append(out, arrReflected.Interface())
    }

	switch nnReflected.Kind() {
	case reflect.Slice :
		for i:= 0; i < nnReflected.Len(); i++ {
			out = append(out,nnReflected.Index(i).Interface())
		}
	default:
		out = append(out, nnReflected.Interface())
	}

	return  out,nil
}

func main() {
	//ar1 := []int{1, 2}
	ar2 := []string{"a","abc"}
	out , err := merge("th",ar2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

}
