package myjson

import (
	"fmt"
	"reflect"
	"strings"
)

const mytag = "json"

/**
 * @api {方法} ./myjson/serialize.go transBool
 * @apiGroup myjson
 * @apiDescription 将bool类型转化为json
 *
 * @apiParam {interface} v bool值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
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

/**
 * @api {方法} ./myjson/serialize.go transString
 * @apiGroup myjson
 * @apiDescription 将string类型转化为json
 *
 * @apiParam {interface} v string值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transString(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])

	str += "\""
	str += reflect.ValueOf(v).String()
	str += "\""

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go transInt
 * @apiGroup myjson
 * @apiDescription 将int类型转化为json
 *
 * @apiParam {interface} v int值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transInt(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])
	var temp = fmt.Sprintf("%v", reflect.ValueOf(v).Int())
	str += temp

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go transUint
 * @apiGroup myjson
 * @apiDescription 将uint类型转化为json
 *
 * @apiParam {interface} v uint值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transUint(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])
	var temp = fmt.Sprintf("%v", reflect.ValueOf(v).Uint())
	str += temp

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go transFloat
 * @apiGroup myjson
 * @apiDescription 将float类型转化为json
 *
 * @apiParam {interface} v float值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transFloat(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])
	var temp = fmt.Sprintf("%v", reflect.ValueOf(v).Float())
	str += temp

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go transMap
 * @apiGroup myjson
 * @apiDescription 将map类型转化为json
 *
 * @apiParam {interface} v map值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transMap(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])
	str += "{"

	for i, key := range reflect.ValueOf(v).MapKeys() {
		if i > 0 {
			str += ","
		}

		str += "\""
		str += key.String()
		str += "\""

		str += ":"

		value, err := JSONMarshal(reflect.ValueOf(v).MapIndex(key).Interface())
		if err != nil {
			panic(err)
		}
		var temp = string(value[:])
		str += temp
	}
	str += "}"

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go transStruct
 * @apiGroup myjson
 * @apiDescription 将struct类型转化为json
 *
 * @apiParam {interface} v struct值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func transStruct(v interface{}, mes []byte) (res []byte) {
	var str = string(mes[:])
	str += "{"

	for i := 0; i < reflect.TypeOf(v).NumField(); i++ {
		fieldName := reflect.ValueOf(v).Type().Field(i).Name

		tag := reflect.ValueOf(v).Type().Field(i).Tag.Get(mytag)

		tags := strings.Split(tag, ",")

		omitemptyContain := strings.Contains(tag, "omitempty")

		if fieldName[0] >= 'a' && fieldName[0] <= 'z' {
			continue
		}
		if tag == "-" {
			continue
		}
		if omitemptyContain && isEmpty(reflect.ValueOf(v).Field(i).Interface()) {
			continue
		}

		if i > 0 {
			str += ","
		}

		if tags[0] != "" {
			fieldName = tags[0]
		}

		str += "\""
		str += fieldName
		str += "\":"

		value, err := JSONMarshal(reflect.ValueOf(v).Field(i).Interface())
		if err != nil {
			panic(err)
		}
		var temp = string(value[:])

		str += temp
	}
	str += "}"

	res = []byte(str)

	return res
}

/**
 * @api {方法} ./myjson/serialize.go isEmpty
 * @apiGroup myjson
 * @apiDescription 将string类型转化为json
 *
 * @apiParam {interface} v string值
 * @apiParam {[]byte} mes 已经格式化输出完成的部分
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 */
func isEmpty(v interface{}) bool {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Bool:
		return value.Bool() == false
	case reflect.String, reflect.Map:
		return value.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	}

	return true
}

/**
 * @api {方法} ./myjson/serialize.go JSONMarshal
 * @apiGroup myjson
 * @apiDescription 将各种类型转化为json，其中支持自定义tag
 *
 * @apiParam {interface} v 参数值
 * @apiSuccess {[]byte} res 将该部分格式化输出后的总信息
 * @apiSuccess {error} err 错误信息
 */
func JSONMarshal(v interface{}) ([]byte, error) {
	var res []byte

	switch reflect.TypeOf(v).Kind() { // 根据接口的类型进行编码
	case reflect.Bool:
		res = transBool(v, res)
	case reflect.String:
		res = transString(v, res)
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		res = transInt(v, res)
	case reflect.Uint, reflect.Uint16, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		res = transUint(v, res)
	case reflect.Float32, reflect.Float64:
		res = transFloat(v, res)
	case reflect.Map:
		res = transMap(v, res)
	case reflect.Struct:
		res = transStruct(v, res)
	}
	//fmt.Println(res)

	return res, nil
}
