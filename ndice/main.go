package main

import (
	"fmt"
	"math"
)

const (
	maxValue = 6
)

// 递归 leetcode超时
func dicesProbability(n int) []float64 {
	if n < 1 {
		return nil
	}
	maxSum := n * maxValue
	res := make([]float64, maxSum-n+1)
	pProbabilities := make([]int, maxSum-n+1)
	Probability(n, pProbabilities)
	total := math.Pow(maxValue, float64(n))
	for i := n; i <= maxSum; i++ {
		ratio := float64(pProbabilities[i-n]) / total
		res[i-n] = ratio
	}
	return res
}

// Probability 不断将点数相加
func Probability(number int, pProbabilities []int) {
	for i := 1; i <= maxValue; i++ {
		ProbabilityHelp(number, number, i, pProbabilities)
	}
}

// ProbabilityHelp 不断将每个骰子的点数相加
func ProbabilityHelp(original int, current int, sum int, pProbabilities []int) {
	if current == 1 {
		// 将点数之和放入数组
		// 如有两个骰子,那么最低数字为2,2-2=0就放入了第一个位置中
		pProbabilities[sum-original]++
	} else {
		for i := 1; i <= maxValue; i++ {
			ProbabilityHelp(original, current-1, sum+i, pProbabilities)
		}
	}
}

func dicesProbability1(n int) []float64 {
	var (
		i, j, k       int
		total         = math.Pow(float64(maxValue), float64(n))
		flag          = 0
		probabilities = make([][]int, 2)
	)
	probabilities[0] = make([]int, n*maxValue+1)
	probabilities[1] = make([]int, n*maxValue+1)

	for i = 1; i <= maxValue; i++ {
		probabilities[flag][i] = 1
	}
	res := make([]float64, 0)
	for k = 2; k <= n; k++ {
		for i = 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}

		for i = k; i <= k*maxValue; i++ {
			probabilities[1-flag][i] = 0
			for j = 1; i-j >= 0 && j <= maxValue; j++ {
				/*
					这一部分的理解，当第n个骰子朝上一面的数为j时(j∈[1,6]),前n-1个骰子朝上一面的点数之和必须为s-j。至于，s-j>=1的这个条件很容易理解，例如当n=2,s=5时，这里的j不能够取5和6，因为第一个骰子不可能骰出0和-1。
					举个例子吧:
					f(2,5) = f(1,1)+f(1,2)+f(1,3)+f(1,4)。
					其中，f(1,1)代表第一次骰1，第二次骰4；
					(1,2)代表第一次骰2,第二次骰3；
					f(1,3)代表第一次骰3，第二次骰2............*/
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}
		flag = 1 - flag
	}
	for i := n; i < maxValue*n; i++ {
		ratio := float64(probabilities[flag][i]) / total
		res = append(res, ratio)
	}
	return res
}

func main() {
	fmt.Println(dicesProbability1(2))
}
