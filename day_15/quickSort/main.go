package main

import (
	"fmt"
)

func quickSort(a []int, left, right int) {
	// set a key
	p := left // key index
	key := a[p]
	i, j := left, right

	fmt.Printf("start left:%d right:%d\n", i+1, j+1)
	fmt.Println(a)
	flag := true
	for i != j {
		if flag { // right
			for j > p {
				if a[j] < key {
					// swap key to a[j]
					fmt.Printf("change %d <-> %d\n", a[j], key)
					temp := a[j]
					a[j] = key
					a[p] = temp
					flag = !flag
					p = j

					fmt.Printf("i:%d j:%d\n", i+1, j+1)
					fmt.Println(a)
					break
				}
				j--
			}
		} else { // left
			for i < p {
				if a[i] > key {
					// swap key to a[i]
					fmt.Printf("change %d <-> %d\n", a[i], key)
					temp := a[i]
					a[i] = key
					a[p] = temp
					flag = !flag
					p = i

					fmt.Printf("i:%d j:%d\n", i+1, j+1)
					fmt.Println(a)
					break
				}
				i++
			}
		}

	}
	// fmt.Println(a)
	// fmt.Println(j)
	fmt.Println("finish a circle")
	if p-1 > left {
		quickSort(a, left, p-1) // left
	}
	if right > p+1 {
		quickSort(a, p+1, right) // right
	}
	// fmt.Println(a[:i])
	// fmt.Println(a[i+1:])
	return
}

func main() {
	var intList = []int{5, 3, 7, 6, 4, 1, 0, 2, 9, 10, 8}
	quickSort(intList, 0, len(intList)-1)
	fmt.Println(intList)
}
