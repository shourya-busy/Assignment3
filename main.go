package main

import (
	"fmt"
	"reflect"
)

//Merges the given interfaces
func merge(arr interface{}, nn interface{}) (interface{},error) {

	//Value is extacted from each interface using reflect
	arrReflected := reflect.ValueOf(arr)
	nnReflected := reflect.ValueOf(nn)
	
	//check for possible invalid inputs
	if arr == nil {
		if nn == nil {
		return nil, fmt.Errorf("Nil Input")
		} else {
			//if 'arr' is nil return 'nn' either as it is or after putting in a slice
			switch nnReflected.Kind() {
			case reflect.Slice:
				return nn,nil
			default :
				return []interface{}{nn},nil
			}
		}
	}

	//Declare an array of interface to store the merged array
	var out []interface{}

	//If both 'arr' and 'nn' are slices then merge them efficiently
	if arrReflected.Kind() == reflect.Slice && nnReflected.Kind() == reflect.Slice {
		out = mergeEfficient(arrReflected,nnReflected,out)
	} else {
		//Else merge them individually
		out = mergeSelect(arrReflected,out)
		out = mergeSelect(nnReflected,out)
	}
	
	//Return the merged array
	return  out,nil
}

func mergeSelect(arr reflect.Value, out []interface{}) []interface{} {
	//Determine the type of 'arr' and put it into out accordingly
	switch arr.Kind() {
	//Push each element of 'arr' if it a slice
	case reflect.Slice :
		for i:= 0; i < arr.Len(); i++ {
			out = append(out,arr.Index(i).Interface())
		}
	default:
	//Push the single element directly 
		out = append(out, arr.Interface())
    }

	return out
}

func mergeEfficient(arrReflected reflect.Value, nnReflected reflect.Value, out []interface{}) []interface{} {
	i := 0

	//Since both are slices---push the values together
	for ; i < arrReflected.Len() && i < nnReflected.Len(); i++ {
		out = append(out,arrReflected.Index(i).Interface())
		out = append(out,nnReflected.Index(i).Interface())
	}

	//Now push the remaining elements if any

	for ; i < arrReflected.Len(); i++ {
		out = append(out,arrReflected.Index(i).Interface())
	}

	for ; i < nnReflected.Len(); i++ {
		out = append(out,nnReflected.Index(i).Interface())
	}

	return out

}

func main() {
	ar1 := []int{1, 2, 3}
	ar2 := []string{"a","abc","efg","123"}

	//function to merge two arrays of any type
	out , err := merge(ar1,ar2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}
