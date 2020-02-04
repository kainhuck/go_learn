package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 快排

func quickSort(A []int, left, right int) {
	// 1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
	i := left
	j := right

	// 2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
	keyIndex := left

	for i < j {
		// 3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]的值交换；
		for j > keyIndex {
			if A[j] < A[keyIndex] {
				swapList(A, j, keyIndex)
				keyIndex = j
				// j--
				break
			}
			j--
		}

		// 4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]的值交换；
		for i < keyIndex {
			if A[i] > A[keyIndex] {
				swapList(A, i, keyIndex)
				keyIndex = i
				// i++
				break
			}
			i++
		}
	}
	// 5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。
	if i-1 > left {
		quickSort(A, left, i-1)
	}
	if i+1 < right {
		quickSort(A, i+1, right)
	}
}

// 交换 intList 的 i 位和 j 位
func swapList(intList []int, i, j int) {
	temp := intList[i]
	intList[i] = intList[j]
	intList[j] = temp
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var intList []int

	for i := 0; i < 10; i++ {
		intList = append(intList, rand.Intn(20))
	}

	fmt.Println("待排序数组:", intList)
	quickSort(intList, 0, len(intList)-1)
	fmt.Println("排完序数组:", intList)
}
