package ttkdescriber

import (
	"fmt"
	"math"
	"sort"

	mathutil "sora.com/math"
)

type Damage struct {
	firerate int
	health   int

	head     int
	headRate int

	neck     int
	neckRate int

	upperTorso     int
	upperTorsoRate int

	lowerTorso     int
	lowerTorsoRate int

	upperArm     int
	upperArmRate int

	lowerArm     int
	lowerArmRate int

	upperLeg     int
	upperLegRate int

	lowerLeg     int
	lowerLegRate int
}

type DamageRate struct {
	location string
	rate     int
}

func GenerateDamageRateMap(damage Damage) map[int]DamageRate {
	drm := make(map[int]DamageRate)

	// 头
	aggregateDamage(drm, damage.head, damage.headRate, "头部")

	// 脖子
	aggregateDamage(drm, damage.neck, damage.neckRate, "脖子")

	// 上胸
	aggregateDamage(drm, damage.upperTorso, damage.upperTorsoRate, "上胸")

	// 腹部
	aggregateDamage(drm, damage.lowerTorso, damage.lowerTorsoRate, "腹部")

	// 上臂
	aggregateDamage(drm, damage.upperArm, damage.upperArmRate, "上臂")

	// 下臂
	aggregateDamage(drm, damage.lowerArm, damage.lowerArmRate, "下臂")

	// 大腿
	aggregateDamage(drm, damage.upperLeg, damage.upperLegRate, "大腿")

	// 小腿
	aggregateDamage(drm, damage.lowerLeg, damage.lowerLegRate, "小腿")

	return drm
}

func aggregateDamage(drm map[int]DamageRate, damage int, rate int, locationName string) {
	if damage > 0 && rate > 0 {
		if v, ok := drm[damage]; ok {
			v.location = v.location + "," + locationName
			v.rate = v.rate + rate
			drm[damage] = v
		} else {
			dr := DamageRate{
				location: locationName,
				rate:     rate,
			}
			drm[damage] = dr
		}
	}
}

func CalCombinationTTK(damage Damage) {
	drm := GenerateDamageRateMap(damage)
	totalRate := 0
	for _, v := range drm {
		totalRate += v.rate
	}

	damages := make([]int, 0)
	for damage := range drm {
		damages = append(damages, damage)
	}

	sort.Ints(damages)

	// 确定所需命中数的边界
	maxDamage := damages[len(damages)-1]
	minCount := int(math.Ceil(float64(damage.health) / float64(maxDamage)))
	minDamage := damages[0]
	maxCount := int(math.Ceil(float64(damage.health) / float64(minDamage)))

	var result [][]int

	for i := minCount; i <= maxCount; i++ {
		damageBacktrack(damage.health, damages, i, []int{}, 0, &result)
	}

	fmt.Printf("达成 %d 血的射击组合有 %d 种：\n", 300, len(result))

	countMap := make(map[int]float64)
	// 倒序遍历结果列表
	for i := len(result) - 1; i >= 0; i-- {
		comb := result[i]
		count := len(comb)
		fmt.Printf("命中方案【%v】，需要【%v】枪致死: \n", len(result)-i, count)
		rate := calRate(comb, drm, totalRate)
		if _, ok := countMap[count]; ok {
			countMap[count] = countMap[count] + rate
		} else {
			countMap[count] = rate
		}
	}

	fmt.Println("在指定命中部位概率的前提下，最终分析:")
	// var curRate float64 = 0
	var keySet []int
	for k, _ := range countMap {
		keySet = append(keySet, k)
	}
	sort.Ints(keySet)
	for i := 0; i < len(keySet); i++ {
		count := keySet[i]
		rate := countMap[keySet[i]]

		displayRate := mathutil.Float64ToStr(rate * 100)
		fmt.Printf("【%v】枪致死的概率为:【%v%%】, TTK:【%v】\n", count, displayRate, ttk(damage.firerate, count))
	}
}

// damageBacktrack 回溯查找所有伤害组合
func damageBacktrack(health int, damages []int, n int, current []int, index int, results *[][]int) {
	if len(current) == n {
		sum := 0
		for _, c := range current {
			sum += c
		}
		if sum >= health {
			combination := make([]int, len(current))
			copy(combination, current)
			*results = append(*results, combination)
		}
		return
	}

	for i := index; i < len(damages); i++ {
		current = append(current, damages[i])
		// 递归
		damageBacktrack(health, damages, n, current, i, results)
		// 回溯
		current = current[:len(current)-1]
	}
}

// 多项式
func calRate(comb []int, drm map[int]DamageRate, totalRate int) float64 {
	n := len(comb)

	var result float64 = float64(mathutil.Factorial(n))

	hitMap := make(map[int]int)
	for i := 0; i < n; i++ {
		damage := comb[i]
		if _, ok := hitMap[damage]; ok {
			hitMap[damage] = hitMap[damage] + 1
		} else {
			hitMap[damage] = 1
		}
	}

	for damage, hit := range hitMap {
		ratio := float64(drm[damage].rate) / float64(totalRate)
		result = result / float64(mathutil.Factorial(hit)) * math.Pow(ratio, float64(hit))
		fmt.Printf("其中【%v】枪【%v】伤害部位【%v】\n", hit, damage, drm[damage].location)
	}

	fmt.Printf("总概率为: %v%%\n", mathutil.Float64ToStr(result*100))
	fmt.Println("--------------------")

	return result

}

// ttk 计算时间（毫秒）
func ttk(rpm int, count int) int {
	ttkValue := 60 * 1000 * (count - 1) / rpm
	return int(ttkValue)
}
