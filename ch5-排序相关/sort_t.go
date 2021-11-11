package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	name  string
	score int
}

// 测试 Search
func TestSliceSearch() {
	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })           // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}
}

type StuScores []Stu

func (s StuScores) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	// sort.Sort(stus)
	// // 判断是否已经排好顺序，将会打印 true
	// fmt.Println("IS Sorted?\n", sort.IsSorted(stus))
	// // 打印排序后的 stus 数据
	// fmt.Println("Sorted:\n", stus)

	// Slice, SliceStable, SliceIsSorted
	sort.SliceStable(stus, func(i, j int) bool { return stus[i].score < stus[j].score })
	fmt.Println("sorted", stus)

	TestSliceSearch()
}
