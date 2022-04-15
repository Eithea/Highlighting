package json

import "encoding/json"


func struct_JSON(inp interface{}) []byte {
	ret, err := json.Marshal(inp)
    if err != nil {
        panic(err)
    }
	return ret
}

type test1 struct {
	Field1 string
	Field2 int
}

func Struct_test1(f1 string, f2 int) []byte {
	st_ret := test1{f1, f2}
	return struct_JSON(st_ret)
}

type test2 struct {
	Field1 string
	Field2 int
}

func Struct_test2(f1 string, f2 int) []byte {
	st_ret := test2{f1, f2}
	return struct_JSON(st_ret)
}