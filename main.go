package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var (
	stringMap = make(map[string]string)
	floatMap  = make(map[string]float64)
	intMap    = make(map[string]int64)
	boolMap   = make(map[string]bool)
)

func main() {

	// str := `{"msg":{"string":"hey","union_type":1}}`
	str := `{"msg":"hey","ts":1232, "active":false, "properties":{"name":"aarjan","address":{"street":"kathmandu"}}}`
	var in map[string]interface{}
	json.Unmarshal([]byte(str), &in)
	o := mapper(in)

	byt, _ := json.MarshalIndent(o, "", " ")
	buf := bytes.NewBuffer(byt)
	buf.WriteTo(os.Stdout)

}

type output struct {
	String map[string]string  `json:"string,omitempty"`
	Int    map[string]int64   `json:"int,omitempty"`
	Float  map[string]float64 `json:"float,omitempty"`
	Bool   map[string]bool    `json:"bool,omitempty"`
	Other  map[string]output  `json:"nested,omitempty"`
}

func mapper(in map[string]interface{}) output {
	out := output{
		String: make(map[string]string, 0),
		Int:    make(map[string]int64, 0),
		Float:  make(map[string]float64, 0),
		Bool:   make(map[string]bool, 0),
		Other:  make(map[string]output, 0),
	}
	for k, v := range in {
		switch t := v.(type) {
		case float64:
			out.Float[k] = v.(float64)
		case int64:
			out.Int[k] = v.(int64)
		case string:
			out.String[k] = v.(string)
		case bool:
			out.Bool[k] = v.(bool)

		case map[string]interface{}:
			out.Other[k] = mapper(v.(map[string]interface{}))
		default:
			fmt.Printf("%T\n", t)
			fmt.Println(k, v)
		}
	}
	return out
}
