package container

import (
	"reflect"
	"log"
)

//////// container lib algorithm
// TODO: package

func Copy(arr interface{}) []interface{} {
	dest := make([]interface{}, 0)

	switch val := arr.(type) {
	case []interface{}:
	case []int:
		for _, v := range val {
			dest = append(dest, v)
		}
	case []string:
		for _, v := range val {
			dest = append(dest, v)
		}
	case int:
		dest = append(dest, val)
	case string:
		dest = append(dest, val)
	}

	return dest
}


func Contains(i interface{}, val interface{}) bool{
	arr := Copy(i)

	if(reflect.TypeOf(arr[0]) != reflect.TypeOf(val)){
		log.Printf("missmatch Type in Contains [%s] and [%s]",reflect.TypeOf(arr[0]) ,reflect.TypeOf(val) )
		panic(0)
	}

	for _, v := range arr{
		if v == val{
			return true
		}
	}
	return false
}
/*
func AllOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v != val{
			return false
		}
	}
	return true
}

func NoneOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v == val{
			return false
		}
	}
	return true
}

func CountIf(arr []Traits, val Traits) int{
	n := 0
	for _, v := range arr{
		if v == val{
			n++
		}
	}
	return n
}

*/
/////////