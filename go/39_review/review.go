/*
 * @Author: xhang 1263403710@qq.com
 * @Date: 2025-04-18 14:03:59
 * @LastEditors: xhang 1263403710@qq.com
 * @LastEditTime: 2025-04-18 14:21:00
 * @FilePath: /github.com/algo/go/39_review/review.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import "fmt"

// result 存储八皇后问题的解，索引表示行，值表示列
var result [8]int

// call8queen 使用回溯法解决八皇后问题
// row: 当前处理的行数
func call8queen(row int) {
	// 递归终止条件：已经放置了8个皇后（处理完最后一行）
	if row == 8 {
		printQueens(result[:]) // 打印当前解
		return
	}

	// 尝试在当前行的每一列放置皇后
	for column := 0; column < 8; column++ {
		// 检查当前位置是否可以放置皇后
		if isOK(row, column) {
			result[row] = column // 在当前位置放置皇后
			call8queen(row + 1)  // 递归处理下一行
		}
	}
}

// isOK 检查在指定位置放置皇后是否合法
// row: 行索引, column: 列索引
// 返回值: 是否可以在该位置放置皇后
func isOK(row int, column int) bool {
	leftup := column - 1  // 左上对角线的起始列位置
	rightup := column + 1 // 右上对角线的起始列位置

	// 逐行向上检查是否有冲突的皇后
	for i := row - 1; i >= 0; i-- {
		// 检查同列是否有皇后
		if result[i] == column {
			return false
		}
		// 检查左上对角线是否有皇后
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
		}
		// 检查右上对角线是否有皇后
		if rightup < 8 {
			if result[i] == rightup {
				return false
			}
		}
		// 更新对角线的列位置
		leftup = leftup - 1
		rightup += 1
	}
	return true // 没有冲突，可以放置皇后
}

// printQueens 打印八皇后问题的一个解
// result: 存储皇后位置的数组，索引是行，值是列
func printQueens(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			// 在皇后位置打印Q，其他位置打印*
			if result[row] == column {
				fmt.Printf("Q")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Println() // 换行，开始打印下一行
	}
	fmt.Println() // 每个解之间空一行
}

// main 程序入口，从第0行开始求解八皇后问题
func main() {
	call8queen(0)
}
