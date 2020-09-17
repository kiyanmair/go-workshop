package main

import "fmt"

func getTypeName(val interface{}) string {
	switch val.(type) {
	case int, int8, int16, int32, int64:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {
	vals := []interface{}{
		1,
		3.14,
		"hello",
		true,
		struct{}{},
	}
	for _, val := range vals {
		typeName := getTypeName(val)
		fmt.Println(val, "is", typeName)
	}
}
