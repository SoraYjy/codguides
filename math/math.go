package math

import (
	"fmt"
	"math"
	"strconv"
)

// 计算组合数 C(a, b)
func Combination(a, b int) float64 {
	if b > a {
		return 0
	}
	return math.Exp(LogFactorial(a) - LogFactorial(b) - LogFactorial(a-b))
}

// 计算对数阶乘，用于计算组合数
func LogFactorial(n int) float64 {
	var (
		sum float64 = 0
		i   int
	)
	for i = 1; i <= n; i++ {
		sum += math.Log(float64(i))
	}
	return sum
}

// 阶乘
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// RoundFloat64 四舍五入float64值到指定的小数位数
func RoundFloat64(f float64, n int) float64 {
	shift := math.Pow(10, float64(n))
	return math.Round(f*shift) / shift
}

// Float64ToStr formats a float64 to a string with two decimal places,
// but keeps leading zeros if there are non-zero digits after them.
func Float64ToStr(f float64) string {
	if f == 0 {
		return "0.00"
	}
	// Convert float64 to string with two decimal places
	str := formatFloat(f, 2)

	var power int = 2

	multipliedFloat := f * 100

	for i := 2; multipliedFloat < 1; i++ {
		power = power + 1
		multipliedFloat = multipliedFloat * 10
		str = formatFloat(f, power)
	}

	return str
}

func formatFloat(f float64, precision int) string {
	return fmt.Sprintf("%."+strconv.Itoa(precision)+"f", f)
}
