/**********
** XML转换
***********/

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant 定义一个结构体
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p *Plant) String() string {
	return fmt.Sprintf("Plant id:%v name:%v origin:%v", p.ID, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 传入我们声明了 XML 的 plant 类型。 使用 MarshalIndent 生成可读性更好的输出结果
	out, _ := xml.MarshalIndent(coffee, "", "")
	fmt.Println(string(out))

	// 明确的为输出结果添加一个通用的 XML 头部信息
	fmt.Println(xml.Header + string(out))

	// 使用 Unmarshal 将 XML 格式的字节流解析到 struct 内。 如果 XML 格式不正确，或无法映射到 struct，将会返回一个描述性错误
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// parent>child>plant 字段标签告诉编码器嵌套 <parent><child>... 下面的所有 plant。
	tomato := &Plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
