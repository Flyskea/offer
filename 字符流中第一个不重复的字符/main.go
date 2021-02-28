package main

import (
	"fmt"
	"math"
)

var occurrence []int = make([]int, 256)
var index int = 1

func Insert(ch byte) {
	if occurrence[ch] >= 1 {
		occurrence[ch] = -1
	} else {
		occurrence[ch] = index
		index++
	}
}

func FirstAppearingOnce() byte {
	var char byte
	minIndex := math.MaxInt64
	for i := 0; i < 256; i++ {
		if occurrence[i] > 0 && occurrence[i] < minIndex {
			char = byte(i)
			minIndex = occurrence[i]
		}
	}
	if char == 0 {
		return '#'
	}
	return char
}

func main() {
	Insert('g')
	fmt.Println(string(FirstAppearingOnce()))
	Insert('o')
	fmt.Println(string(FirstAppearingOnce()))
	Insert('o')
	fmt.Println(string(FirstAppearingOnce()))
	Insert('g')
	fmt.Println(string(FirstAppearingOnce()))
	Insert('l')
	fmt.Println(string(FirstAppearingOnce()))
	Insert('e')
	fmt.Println(string(FirstAppearingOnce()))
}
