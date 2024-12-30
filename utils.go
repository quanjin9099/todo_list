package main

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// TodoInfoSlice 按照 TodoInfo.Score 从大到小排序
type TodoInfoSlice []TodoInfo

func (a TodoInfoSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a TodoInfoSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a TodoInfoSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Score < a[i].Score
}
