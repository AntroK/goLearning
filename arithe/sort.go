package main

import "fmt"

func partition(data []int, first int, end int) int {
	var temp = 0
	var i = first
	var j = end
	for {
		for j > i{
			if data[j] < data[i]{
				temp = data[j]
				data[j] = data[i]
				data[i] = temp
				break
			}
			j --
		}
		for i < j{
			if data[j] < data[i]{
				temp = data[j]
				data[j] = data[i]
				data[i] = temp
				break
			}
			i ++
		}
		if i >= j {
			break
		}
	}
	return i
}

// 递归函数
func quickSort(data []int, low int, high int) {
	var pivotPos int
	if low < high {
		// 找位置
		pivotPos = partition(data, low, high)
		//左右递归
		quickSort(data, low, pivotPos-1)
		quickSort(data, pivotPos+1, high)
	}
}

func quick(data []int, first int, end int) int {
	var temp = 0
	var i = first
	var j = end
	for {
		for j > i{
			if data[j] < data[i]{
				temp = data[j]
				data[j] = data[i]
				data[i] = temp
				break
			}
			j --
		}
		for i < j{
			if data[j] < data[i]{
				temp = data[j]
				data[j] = data[i]
				data[i] = temp
				break
			}
			i ++
		}
		if i >= j {
			break
		}
	}

	return i
}

func twoSum(nums []int, target int) []int {
	var datas = make(map[int]int)
	for k, v:=range nums{
		datas[v] = k
	}

	for k, v:=range datas{
		if datas[target - k] != 0{
			return []int{v, datas[target-k]}
		}
	}
	return []int{}
}

func main(){
	var a = []int{1,4,3,3}
	quickSort(a, 0, 3)
	fmt.Println(twoSum(a, 6))
}
