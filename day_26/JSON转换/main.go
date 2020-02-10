/**********
** Json转换
***********/

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kainhuck/fansylog"
)

var log = fansylog.NewConsoleLogger("debug")

// 定义两个结构体
// 只有 可导出 的字段才会被 JSON 编码/解码。必须以大写字母开头的字段才是可导出的
type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 首先我们来看一下基本数据类型到 JSON 字符串的编码过程。 这是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	flt, _ := json.Marshal(1.23)
	fmt.Println(string(flt))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这是一些切片和 map 编码成 JSON 数组和对象的例子
	slcD := []string{"huhu", "love", "tutu"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"tutu": 605, "huhu": 1022}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear", "草莓"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据的键名。 上面 Response2 的定义，就是这种标签的一个例子
	res2D := &response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear", "草莓"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//=====================================================================
	//== 现在来看看将 JSON 数据解码为 Go 值的过程。 这是一个普通数据结构的解码例子。==
	//=====================================================================

	byt := []byte(`{"name":"tutu", "love":["pingpong", "huhu"], "age":1.12}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		// panic(err)
		log.Fatal("error")
	}
	age := dat["age"].(float64)
	fmt.Println(age)

	// 访问嵌套的值需要一系列的转化
	love := dat["love"].([]interface{})
	me := love[1].(string)
	fmt.Println(me)

	// 我们还可以将 JSON 解码为自定义数据类型。 这样做的好处是，可以为我们的程序增加附加的类型安全性， 并在访问解码后的数据时不需要类型断言。
	str := `{"page":2,"fruits":["apple","peach","pear","草莓"]}`
	var res2 response2
	if err := json.Unmarshal([]byte(str), &res2); err != nil {
		// panic(err)
		log.Fatal("error")
	}
	fmt.Println(res2.Fruits[3])

	// 在上面例子的标准输出上， 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。 当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
