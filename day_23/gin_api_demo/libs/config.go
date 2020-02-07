/******************************
** 解析 config.ini 文件
*******************************/

package libs

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// Configs ...
type Configs struct {
	config map[string]string
	node   string
}

// Conf ...
var Conf *Configs

func init() {
	Conf = new(Configs)
	Conf.LoadConfig("config/config.ini")
}

// LoadConfig ...
func (conf *Configs) LoadConfig(path string) {
	// 1. 打开文件获取文件对象
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
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

		

	}

}

// Read ...
func (conf *Configs) Read(node, key string) string {
	return "nil"
}
