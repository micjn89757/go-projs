package test

import (
	"lottery/util"
	"math/rand"
	"slices"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	const L = 100

	for c := 0; c < 30; c++ {
		arr := make([]float64, 0, L)
		for i := 0; i < L; i++ {
			arr = append(arr, rand.Float64())
		}

		slices.Sort(arr) // 排序概率

		var target float64

		// 测试两个越界情况
		target = arr[0] - 1.0
		if util.BinarySearch(arr, target) != 0 {
			t.Fail()
		}

		target = arr[len(arr)-1] + 1.0
		if util.BinarySearch(arr, target) != 0 {
			t.Fail()
		}

		// 每个分割点以及2个分割点中间的值都要测试
		target = arr[0]
		if util.BinarySearch(arr, target) != 0 {
			t.Fail()
		}

		for i := 0; i < len(arr)-1; i++ {
			target = (arr[i] + arr[i+1]) / 2

			if util.BinarySearch(arr, target) != i+1 {
				t.Fail()
			}

			target = arr[i+1]

			if util.BinarySearch(arr, target) != i+1 {
				t.Fail()
			}
		}

	}
}

// 测试按照比例抽奖
func TestLottery(t *testing.T) {
	probs := []float64{5, 2, 4}                   // 指定抽中概率
	countMap := make(map[int]float64, len(probs)) // 记录用户抽中奖品个数

	// 模拟10000次抽奖
	for i := 0; i < 10000; i++ {
		index := util.Lottery(probs)
		countMap[index] += 1
	}

	// 以下值应当很接近
	t.Log(countMap[0] / probs[0])
	t.Log(countMap[1] / probs[1])
	t.Log(countMap[2] / probs[2])
}

// go test -v ./util/test -run=^TestBinarySearch$ -count=1
// go test -v ./util/test -run=^TestLottery$ -count=1
