package json

import "encoding/json"


func struct_JSON(inp interface{}) []byte {
	ret, err := json.Marshal(inp)
    if err != nil {
        panic(err)
    }
	return ret
}

type testdata struct {
	Field1 string
	Field2 int
}

func Struct_testdata(f1 string, f2 int) []byte {
	st_ret := testdata{f1, f2}
	return struct_JSON(st_ret)
}