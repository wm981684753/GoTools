package main

import "fmt"

func main() {
	// targetArr := []int{1, 7, 3, 6, 5, 6}
	// targetArr := []int{1, 2, 3}
	targetArr := []int{2, 1, -1}
	// targetArr := []int{1, -1, 2}
	midIndex := -1
	for i, _ := range targetArr {
		midIndex = i
		leftV := 0
		rightV := 0
		for j, v := range targetArr {
			if j < i {
				leftV += v
			}
			if j > i {
				rightV += v
			}
		}
		if leftV == rightV {
			fmt.Println(midIndex)
			break
		}
	}
	if midIndex == -1 {
		fmt.Println(midIndex)
	}
}
