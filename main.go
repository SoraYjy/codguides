package main

import (
	"fmt"
)

// coinCombinations 用回溯算法生成所有硬币组合
func coinCombinations(coins []int, n int, current []int, index int, results *[][]int) {
	if len(current) == n { // 如果已经选了 n 个硬币
		sum := 0
		for _, c := range current {
			sum += c
		}
		if sum > 300 { // 筛选面值总和 > 300 的组合
			combination := make([]int, len(current))
			copy(combination, current)
			*results = append(*results, combination)
		}
		return
	}

	// 遍历硬币面值，从 index 开始避免重复排列
	for i := index; i < len(coins); i++ {
		current = append(current, coins[i])             // 选择当前硬币
		coinCombinations(coins, n, current, i, results) // 递归处理
		current = current[:len(current)-1]              // 回溯
	}
}

func main() {
	// 定义硬币面值
	coins := []int{50, 100, 200, 300}

	// 定义需要选择的硬币数量
	n := 4

	// 存储结果的切片
	var results [][]int

	// 生成组合
	coinCombinations(coins, n, []int{}, 0, &results)

	// 输出结果
	fmt.Println("满足条件的硬币组合（总面值 > 300）：")
	for _, combination := range results {
		fmt.Println(combination)
	}
}
