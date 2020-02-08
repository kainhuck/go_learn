/******************************
** 解析 config.ini 文件
*******************************/

package libs

import (
	"bufio"
	"go_learn/day_10/mylogger"
	"io"
	"os"
	"strings"
)

// Configs ...
type Configs struct {
	config map[string]string
	node   string
}

var (
	// Conf ...
	Conf *Configs

	// MidStr ...
	MidStr = "<@_@>"

	log = mylogger.NewConsoleLogger("Debug")
)

func init() {
	Conf = new(Configs)
	Conf.LoadConfig("config/config.ini")
}

// LoadConfig ...
func (conf *Configs) LoadConfig(path string) {
	conf.config = make(map[string]string)
	// 1. 打开文件获取文件对象
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		lines, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(lines)) // 前后去空格
		if err != nil {
			if err == io.EOF {
				break // 读完最后一行,退出
			}
			log.Fatal(err.Error())
		}

		// 处理注释
		if strings.Index(line, "#") == 0 {
			continue
		}

		// [xxx]
		n := strings.Index(line, "[")
		nl := strings.Index(line, "]")

		if n > -1 && nl > -1 && nl > n+1 {
			conf.node = strings.TrimSpace(line[n+1 : nl])
			continue
		}

		if len(conf.node) == 0 || len(line) == 0 {
			continue
		}

		arr := strings.Split(line, "=")
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		newKey := conf.node + MidStr + key
		conf.config[newKey] = value
	}

}

// Read ...
func (conf *Configs) Read(node, key string) string {
	newKey := node + MidStr + key
	value, ok := conf.config[newKey]
	if !ok {
		return ""
	}
	return value

}
