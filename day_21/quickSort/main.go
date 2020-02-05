package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var a []int

	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}

	fmt.Printf("未排序之前的数组: %v\n", a)
	quickSort(a, 0, len(a)-1)
	fmt.Printf("排完序之后的数组: %v\n", a)
}

func quickSort(a []int, left, right int) {
	// 1. 设置 i = left, j = right
	i, j := left, right

	// 2. 设置 key = left
	key := left

	// 5. 重复上述(3,4)操作,直到 i >= j
	for i < j {
		// 3. 从后向前查找,找到比key对应的值小的就交换顺序
		for j > key {
			if a[j] < a[key] {
				swap(a, j, key)
				key = j
				break
			}
			j--
		}

		// 4. 从前向后查找,找到比key对应值大的就交换顺序
		for i < key {
			if a[i] > a[key] {
				swap(a, i, key)
				key = i
				break
			}
			i++
		}
	}

	// 递归
	if i-1 > left {
		quickSort(a, left, i-1) // 左边
	}

	if right > i+1 {
		quickSort(a, i+1, right) // 右边
	}

}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
