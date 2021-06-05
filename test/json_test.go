package test

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {
	jstr := []byte(`{"1":{"a":1, "b":2}, "2":{"a":3, "b":4}}`)
	var data map[string]interface{} // map[1] = interface{} map[2]=interface{}
	json.Unmarshal(jstr, &data)

	t.Log(data)
}

func TestJson_Print(t *testing.T) {
	type Person struct {
		Name string `json:"object"`
		Age  int
	}

	p := Person{}

	json.Unmarshal([]byte(`{"object":"tom","Age":10}`), &p)
	t.Log(p)
}
