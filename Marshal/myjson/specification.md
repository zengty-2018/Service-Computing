# JSONMarshal函数设计说明
## 对传入参数interface{}的处理
使用reflect包进行接口信息的判断。<br>
使用reflect.Typeof()获取该数据的类型。<BR>
使用reflect.Valueof()获取该数据的值，但是这样获取的值不能够直接使用，要再根据对应类型转为能直接使用的值，如：reflect.Valueof(v).Int()。<BR>
<BR>
由于直接对byte数组的操作比较困难，因此在本包中将全部byte数组转为string操作，在返回时才再次转回byte数组。<BR>

对于bool、int、float等基本数据类型，直接调用相应的transBool()、transInt()等函数，这里以transBool()的代码为例展示：
```
func transBool(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])

	s := reflect.ValueOf(v)

	if s.Bool() {
		str += "true"
	} else {
		str += "false"
	}

	res = []byte(str)

	return res
}
```
对于map，使用reflect.Valueof(v).MapKeys()遍历整个map，然后对每一个key映射的对象递归调用JSONMarshal。<br>
对于struct，总体与map类似，主要是增加了标签内容，对于每一个field，使用file.Tag.Get(mytag)方法获取标签，根据json的规则，filename首字母为小写的是私有变量不能读取，tag为'-'的是强制跳过不进行解析的标志；tag中包含omitempty的是如果该字段为空则不进行解析的标志，因此还需要增加一个isEmpty函数判断字段是否为空。源码较长，可以到serialize.go中查阅。

## 判断接口类型并转化为json格式
本包中支持的类型有：
- bool类型：直接转为true或false；
- string类型：在字符串两边加上双引号；
- int、uint、float类型：直接转化为字符串形式；
- map类型：两边加上花括号，以key:value形式写入，对与对之间用逗号分隔；
- struct：两边加上花括号，以name:JSONMarshal(filed)形式递归调用写入。

## 测试
测试的想法是将自定义标签改为json，然后调用标准encoding/json中的marshal函数，与本包的结果进行对比。
```
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
```

结果：
```
PS D:\Service Computing\Marshal\myjson> go test
{"A":true,"B":"bbbbbb","C":10,"D":2,"E":33,"F":{"ee":111,"ww":222}}
{"A":true,"B":"bbbbbb","C":10,"D":2,"E":33,"F":{"ee":111,"ww":222}}
{"name":false}
{"name":false}
PASS
ok      _/D_/Service_Computing/Marshal/myjson   0.442s
```