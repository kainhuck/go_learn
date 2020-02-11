package main

import (
	"crypto/sha1"
	"fmt"
)

/*
** SHA1 散列（hash） 经常用于生成二进制文件或者文本块的短标识。
 */

func main() {
	// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 然后 sha1.Sum([]byte{})
	s := "sha1 this string"
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串， 需要使用 []byte(s) 将其强制转换成字节数组
	h.Write([]byte(s))

	// Sum 得到最终的散列值的字符切片。Sum 接收一个参数， 可以用来给现有的字符切片追加额外的字节切片：但是一般都不需要这样做
	bs := h.Sum(nil)

	fmt.Println(s)
	// SHA1 值经常以 16 进制输出，例如在 git commit 中。 我们这里也使用 %x 来将散列结果格式化为 16 进制字符串
	fmt.Printf("%x\n", bs)

	// 你可以使用和上面相似的方式来计算其他形式的散列值。 例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法
	// h := md5.New()
	// h.Write([]byte("I Love xsy"))
	// bs := h.Sum(nil)
	// fmt.Printf("%x\n", bs)
}
