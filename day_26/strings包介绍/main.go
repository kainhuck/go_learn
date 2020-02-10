/**************************
** strings 包的常用函数介绍
** 以及别的strings操作
***************************/
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("COunt:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"tutu", "kangkang"}, "&"))
	p("Repeat:", s.Repeat("tutu", 9))
	p("Replace:", s.Replace("test", "t", "h", -1))
	p("Replace:", s.Replace("test", "t", "h", 1))
	p("Split:", s.Split("test", "s"))
	p("ToLower:", s.ToLower("TEsT"))
	p("ToUpper:", s.ToUpper("teSt"))
	p()

	p("Len:", len("test"))
	p("Char:", "test"[1])
}
