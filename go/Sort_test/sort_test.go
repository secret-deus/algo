package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateRandomArray(n int) []int {
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano()) // 种子
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(10000) // 随机数范围 0-9999
	}
	return arr
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer() // 重置计时器
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr)) // 复制数组
		copy(copyArr, arr)
		BubbleSort(copyArr)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		InsertionSort(copyArr)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		SelectionSort(copyArr)
	}
}

func BenchmarkShellSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		ShellSort(copyArr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		_ = MergeSort(copyArr) // 注意：MergeSort 返回新数组
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		QuickSort(copyArr)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	arr := generateRandomArray(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyArr := make([]int, len(arr))
		copy(copyArr, arr)
		HeapSort(copyArr)
	}
} // 确保所有函数都有闭合的大括号

// 文件最后需要换行符
