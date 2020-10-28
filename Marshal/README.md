# 服务计算第五次作业
## 任务
实现encoding/json中Marshal函数类似功能的函数，将各种数据结构的数据转化成json字符流。

## Marshal函数功能介绍
函数通过传入一个数据的接口v interface{}，将其转为json字符流。<br>
- 如果是int、bool、string等基本数据类型，则直接转为对应字符。<br>
- 如果是map、struct等复杂结构，则会通过递归调用的方式将其转为json数组的形式。<br>

除此之外，包还提供了json中tag的功能：
- 当标签中有'-'时，不解析该字段
- 当标签中有'omitempty'时，当该字段为空时不解析
- 当标签中有'filename'时，解析是使用这个名称

默认tag为json，可以通过修改serialize.go中的mytag来改变。

测试样例：(为了测试自定义tag，这里将mytag改为了tag）<BR>
```
type test1 struct {
	A bool
	B string
	C uint
	D int
	E int32
	F map[string]int
}
t1 := test1{
    A: true,
    B: "bbbbbb",
    C: 10,
    D: 2,
    E: 33,
    F: map[string]int{
        "qq": 111,
        "ww": 222,
    },
}

type test2 struct {
	A bool   `tag:"name"`
	B string `tag:"-"`
	C int    `tag:"docs,omitempty"`
}
t2 := test2{
    A: false,
    B: "bbbbb",
}
```
输出：
```
t1:{"A":true,"B":"bbbbbb","C":10,"D":2,"E":33,"F":{"qq":111,"ww":222}}
t2:{"name":false}
```
## 生成api文档
详情见api中的html文件。