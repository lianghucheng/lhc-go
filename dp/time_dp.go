package main

import (
	"fmt"
)

var mem [][]int
var pix int

func main() {
	N:= 10
	T:= 2
	Li := []int{2,3,1,5,7,3,4,3,1,30}
	Ei:= []int{3,2,5,2,6,2,2,2,8,1}
	x := make([][]int, N)
	for k:=range x {
		x[k] = make([]int, N+2)
	}
	mem = x
	//fmt.Scanf("%d%d",&N,&T)
	//Li :=make([]int, N)
	//Ei:= make([]int, N)
	//for i:= 0;i<N;i++ {
	//	val1,val2 := 0,0
	//	fmt.Scanf("%d%d",&val1,&val2)
	//	Li[i]= val1
	//	Ei[i]=val2
	//}

	fmt.Println(solve1(Li,Ei,T,N,N-1,0))
}

func res(Li []int, d int){
	for i:=0;i<len(Li);i++ {
		if Li[i]>0{
			Li[i] -= d
		}
	}
}

func solve1(Li []int,Ei []int,T,N,idx int,tm int) (int) {
	fmt.Println("solve1",tm,idx)
	if idx < 0 {
		return 0
	}
	if mem[pix][idx+2] != 0 {
		return mem[pix][idx+2]
	}
	value1 := 0
	if Li[idx]>0 {
		value1 =Ei[idx]
		Lt := append([]int{},Li...)
		Lt[idx] =-1
		res(Lt,T)
		val := solve2(Lt,Ei,T,N, tm + 1)
		value1 += val
	}

	value2 := 0
	value2 = solve1(Li, Ei,T,N,idx-1,tm)

	max := 0
	if value1 < value2 {
		max = value2
	} else {
		max =value1
	}

	fmt.Println("result solve1",max, idx)
	mem[pix][idx+2] = max
	return max
}

func solve2(Li []int,Ei []int,T,N int,tm int) (int) {
	pix = tm
	idx := -1
	for i:=len(Li)-1;i>=0;i-- {
		if Li[i]>0 {
			idx = i
		}
	}
	fmt.Println("solve2",tm,idx)

	value1 := 0
	if idx >=0 {
		if Li[idx]>0 {
			value1 =Ei[idx]
			Lt := append([]int{},Li...)
			Lt[idx] =-1
			res(Lt,T)
			val := solve2(Lt,Ei,T,N, tm + 1)
			value1 += val
		}
	}
	value2 := 0
	value2 = solve1(Li, Ei,T,N,idx-1,tm)


	max := 0
	if value1 < value2 {
		max = value2
	} else {
		max =value1
	}
	fmt.Println("result solve2",max, idx)
	return max
}

//深度和最后一个idx