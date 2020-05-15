package main

import (
	"fmt"
	"math"
)

var result = make([][]int, 0)

func main() {
	w:=[]int{1,23,4,6,5}
	v:=[]int{1}
	s:=1
	answer:=solve(w,v,s,len(w)-1)
	fmt.Println(answer)
}

func solve(w,v []int, s int, index int) int {
	if index < 0 {
		return 0
	}
	if result[index][s] == 0{
		result[index][s] =int(math.Max(float64(v[index] + solve(w,v,s-w[index],index-1)),float64(solve(w,v,s,index-1))))
	}
	return result[index][s]
}