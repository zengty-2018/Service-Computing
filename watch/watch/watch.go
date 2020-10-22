package watch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var comment byte

/**
 *@api {方法} ./watch/watch.go Init
 *@apiGroup watch
 *@apiDescription 调用runtime.GOOS判断系统类型，从而确定注释符
 *
 */

func Init() {
	if runtime.GOOS == "linux" {
		comment = '#'
	} else if runtime.GOOS == "windows" {
		comment = ';'
	}

	//fmt.Println(runtime.GOOS)
}

//签名
type ListenFunc func(string)

//接口
type Listener interface {
	listen(file_name string)
}

/**
 * @api {方法} ./watch/watch.go listen
 * @apiGroup watch
 * @apiDescription 监听文件是否被修改的函数
 *
 * @apiParam {string} file_name 文件名
 */
func (Listen ListenFunc) listen(file_name string) {
	start_file_mes := getFileMes(file_name)
	flag := false

	for {
		cur_file_mes := getFileMes(file_name)

		len1 := len(start_file_mes)
		len2 := len(cur_file_mes)

		if len1 != len2 {
			break
		}

		for i := 0; i < len1; i++ {
			if start_file_mes[i] != cur_file_mes[i] {
				flag = true
				break
			}
		}
		//上面都是判断是否被修改（与一开始的字段是否相等）

		if flag {
			fmt.Println("The file have been changed.")
			break
		}

		//time.Sleep(1000000000)
	}
}

/**
 * @api {方法} ./watch/watch.go getFileMes
 * @apiGroup watch
 * @apiDescription 获取文件内容
 *
 * @apiParam {string} file_name 文件名
 * @apiSuccess {map[string]string} file_mes 文件内容
 */
func getFileMes(file_name string) []string {
	var file_mes []string
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Get file message error!")
		return file_mes
	}

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		file_mes = append(file_mes, str)
	}

	file.Close()
	return file_mes
}

//自定义错误类型
type error_mes struct {
	mes   string
	type_ int
}

//得到错误信息
func GetMes(err error_mes) string {
	return err.mes
}

func get_new_error(mes string, type_ int) error_mes {
	return error_mes{mes, type_}
}

func (err *error_mes) out_error() string {
	fmt.Println(err.mes)
	return err.mes
}

/**
 * @api {方法} ./watch/watch.go read_file
 * @apiGroup watch
 * @apiDescription 读取文件并返回map[string]string和自定义错误类型
 *
 * @apiParam {string} file_name 文件名
 * @apiSuccess {map[string]string} file_mes 文件内容
 * @apiSuccess {error_mes} err 自定义错误信息
 */
func Read_file(file_name string) (map[string]string, error_mes) {
	var error error_mes
	file, err := os.Open(file_name)
	configuration := make(map[string]string)

	if err != nil {
		error = get_new_error("open ili file error.", 0)
		return configuration, error
	} else {
		error = get_new_error("", 1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		//逐行读取
		str, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		//删去空格
		str = strings.TrimSpace(str)

		//如果是空的/注释/则忽略
		if str == "" || str[0] == '[' || str[0] == comment {
			continue
		}

		pair := strings.Split(str, "=")

		if len(pair) == 2 {
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])
			configuration[key] = value
			// fmt.Println(key)
		}

		if err != nil {
			// fmt.Println(err)
			break
		}
	}

	return configuration, error
}

/**
 * @api {方法} ./watch/watch.go Watch
 * @apiGroup watch
 * @apiDescription 读取、打印、监听文件，如果文件内容被修改则输出修改后的文件内容
 *
 * @apiParam {string} file_name 文件名
 * @apiParam {Listener} Listen 监听器
 * @apiSuccess {map[string]string} file_mes 文件内容
 * @apiSuccess {error_mes} err 自定义错误信息
 */
func Watch(file_name string, Listen Listener) (map[string]string, error_mes) {
	Init()

	config, err := Read_file(file_name)

	if len(err.mes) != 0 {
		fmt.Println(err.mes)
		return config, err
	}
	fmt.Println("Listening....")
	Print_config(config)
	Listen.listen(file_name)
	fmt.Println("The file have been changed.")
	fmt.Println("")

	config, err = Read_file(file_name)
	if len(err.mes) != 0 {
		fmt.Println(err.mes)
		return config, err
	}
	//Print_config(config)

	return config, err
}

/**
 * @api {方法} ./watch/watch.go print_config
 * @apiGroup watch
 * @apiDescription 输出相应config
 *
 * @apiParam {map[string]string} config 保存配置文件的对应关系
 */
func Print_config(config map[string]string) {
	fmt.Println("Configure:")
	for key, value := range config {
		fmt.Println(key, "=", value)
		time.Sleep(1500000000)
	}
	fmt.Println("Print configure end.")

}
