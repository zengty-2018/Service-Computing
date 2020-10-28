package myjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type test1 struct {
	A bool
	B string
	C uint
	D int
	E int32
	F map[string]int
}

type test2 struct {
	A bool   `json:"name"`
	B string `json:"-"`
	C int    `json:"docs,omitempty"`
}

func TestJSONMarshal(t *testing.T) {
	t1 := test1{
		A: true,
		B: "bbbbbb",
		C: 10,
		D: 2,
		E: 33,
		F: map[string]int{
			"ee": 111,
			"ww": 222,
		},
	}

	get, err1 := JSONMarshal(t1)
	if err1 != nil {
		panic(err1)
	}
	want, err2 := json.Marshal(t1)
	if err2 != nil {
		panic(err2)
	}
	var get1 = string(get[:])
	var want1 = string(want[:])

	fmt.Println(get1)
	fmt.Println(want1)

	if !reflect.DeepEqual(get, want) {
		t.Error("want : " + want1 + "\n" + "get : " + get1)
	}

}

func TestJSONMarshalTag(t *testing.T) {
	t2 := test2{
		A: false,
		B: "bbbbb",
	}

	get, err1 := JSONMarshal(t2)
	if err1 != nil {
		panic(err1)
	}
	want, err2 := json.Marshal(t2)
	if err2 != nil {
		panic(err2)
	}

	var get2 = string(get[:])
	var want2 = string(want[:])

	fmt.Println(get2)
	fmt.Println(want2)

	if !reflect.DeepEqual(get, want) {
		t.Error("want : " + want2 + "\n" + "get : " + get2)
	}
}
