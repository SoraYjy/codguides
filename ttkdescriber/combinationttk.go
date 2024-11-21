package main

import (
	"fmt"
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

func NewDamage() Damage {
	damage := Damage{}
	damage.firerate = 800
	damage.health = 300

	damage.headRate = 15
	damage.neckRate = 5
	damage.upperTorsoRate = 30
	damage.lowerTorsoRate = 20
	damage.upperArmRate = 10
	damage.lowerArmRate = 5
	damage.upperLegRate = 10
	damage.lowerLegRate = 5

	damage.head = 38
	damage.neck = 35
	damage.upperTorso = 35
	damage.lowerTorso = 35
	damage.upperArm = 35
	damage.lowerArm = 35
	damage.upperLeg = 31
	damage.lowerLeg = 31

	return damage
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
	if damage > 0 {
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

func main() {
	damage := NewDamage()
	calCombinationTTK(damage)
}

func calCombinationTTK(damage Damage) {
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

	var result [][]int
	var combinations []int

	damgeBacktrack(damages, 300, combinations, &result, 0)

	fmt.Printf("达成 %d 血的射击组合有 %d 种：\n", 300, len(result))

	countMap := make(map[int]float64)
	// 倒序遍历结果列表
	for i := len(result) - 1; i >= 0; i-- {
		comb := result[i]
		count := len(comb)
		fmt.Printf("命中方案【%v】，需要【%v】枪致死: \n", i+1, count)
		rate := calRate(comb, drm, totalRate)
		if _, ok := countMap[count]; ok {
			countMap[count] = countMap[count] + rate
		} else {
			countMap[count] = rate
		}
	}

	fmt.Println("在指定命中部位概率的前提下，最终分析:")
	var curRate float64 = 0
	var keySet []int
	for k, _ := range countMap {
		keySet = append(keySet, k)
	}
	for i := 0; i < len(keySet); i++ {
		count := keySet[i]

		if i == len(keySet)-1 {
			// 处理边界问题，最多所需子弹数下，100%击杀
			curRate = 1
		} else {
			rate := countMap[keySet[i]]
			curRate = (1 - curRate) * rate
		}
		displayRate := mathutil.RoundFloat64(curRate, 7) * 100
		fmt.Printf("【%v】枪致死的概率为:【%v%%】, TTK:【%v】\n", count, displayRate, ttk(damage.firerate, count))
	}

}

// backtrack 回溯找所有伤害组合
func damgeBacktrack(damages []int, target int, combination []int, result *[][]int, start int) {
	if target <= 0 {
		// 如果总伤害量达标，则加入结果集
		comb := make([]int, len(combination))
		copy(comb, combination)
		*result = append(*result, comb)
		return
	}
	for i := start; i < len(damages); i++ {
		// 加入当前伤害量，并递归调用
		combination = append(combination, damages[i])
		damgeBacktrack(damages, target-damages[i], combination, result, i)
		// 回溯，移除最后一个伤害值
		combination = combination[:len(combination)-1]
	}
}

func calRate(comb []int, drm map[int]DamageRate, totalRate int) float64 {
	hitMap := make(map[int]int)

	for _, damage := range comb {
		if _, ok := hitMap[damage]; ok {
			hitMap[damage] = hitMap[damage] + 1
		} else {
			hitMap[damage] = 1
		}
	}

	var hitComb float64 = 1
	var totalComb float64
	var total int = 0
	for damage, hitCount := range hitMap {
		hitComb = hitComb * mathutil.Combination(drm[damage].rate, hitCount)
		total = total + drm[damage].rate
		fmt.Printf("其中【%v】枪【%v】伤害部位【%v】\n", hitCount, damage, drm[damage].location)
	}

	totalComb = hitComb / mathutil.Combination(totalRate, len(comb))

	fmt.Printf("总概率为: %v\n", totalComb)
	fmt.Println("--------------------")

	return totalComb
}

// ttk 计算时间（毫秒）
func ttk(rpm int, count int) int {
	// 计算ttk
	ttkValue := 60 * 1000 * (count - 1) / rpm
	return int(ttkValue)
}
