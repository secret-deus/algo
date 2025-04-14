package main

// BubbleSort 冒泡排序：通过反复交换相邻元素实现排序。
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
// 适用于小型数据集，但效率较低。
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // 优化：如果没有交换，提前结束
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// InsertionSort 插入排序：将元素逐一插入已排序的子数组中。
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
// 适合近乎有序的数据。
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // 当前元素
		j := i - 1    // 已排序子数组的最后一个索引
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // 向右移动
			j--
		}
		arr[j+1] = key // 插入元素
	}
}

// SelectionSort 选择排序：每次找到最小元素并交换到正确位置。
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
// 简单但不适合大型数据集。
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j // 找到最小元素的索引
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i] // 交换
		}
	}
}

// ShellSort 希尔排序：基于插入排序的优化，通过增量分组。
// 时间复杂度：O(n log n) (平均)
// 空间复杂度：O(1)
// 适合中等规模数据。
func ShellSort(arr []int) {
	n := len(arr)
	gap := n / 2 // 初始增量
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2 // 减小增量
	}
}

// MergeSort 归并排序：分治策略，先递归拆分数组，然后合并。
// 时间复杂度：O(n log n)
// 空间复杂度：O(n) (辅助空间)
// 稳定且高效，适合大型数据集。
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])  // 递归左半部分
	right := MergeSort(arr[mid:]) // 递归右半部分
	return merge(left, right)     // 合并
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)  // 剩余元素
	result = append(result, right[j:]...) // 剩余元素
	return result
}

// QuickSort 快速排序：分治策略，选择基准元素分区。
// 时间复杂度：O(n log n) (平均)，O(n²) (最坏)
// 空间复杂度：O(log n) (递归栈)
// 高效但可能退化。
func QuickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)  // 分区
		quickSortHelper(arr, low, pi-1)  // 递归左部分
		quickSortHelper(arr, pi+1, high) // 递归右部分
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // 选择最后一个作为基准
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // 交换
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // 放置基准
	return i + 1
}

// HeapSort 堆排序：构建堆，然后排序。
// 时间复杂度：O(n log n)
// 空间复杂度：O(1)
// 适合大型数据集。
func HeapSort(arr []int) {
	n := len(arr)
	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// 提取元素
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换根和最后一个
		heapify(arr, i, 0)              // 重新堆化
	}
}

func heapify(arr []int, n int, i int) {
	largest := i     // 根作为最大
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换
		heapify(arr, n, largest)                    // 递归
	}
}
